package tasks

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/eskpil/rockferry/pkg/mac"
	"github.com/eskpil/rockferry/pkg/rockferry"
	"github.com/eskpil/rockferry/pkg/rockferry/spec"
	"github.com/google/uuid"
	"github.com/siderolabs/go-pointer"
)

func ensurePath(p string) error {
	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func downloadFile(dst string, remote string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(remote)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

type CreateVirtualMachineTask struct {
	Request *rockferry.MachineRequest
}

func (t *CreateVirtualMachineTask) createVmDisks(ctx context.Context, executor *Executor) ([]*spec.MachineSpecDisk, error) {
	disks := []*spec.MachineSpecDisk{}

	for _, disk := range t.Request.Spec.Disks {
		pool, err := executor.Rockferry.StoragePools().Get(ctx, disk.Pool, nil)
		if err != nil {
			return nil, err
		}

		d := new(spec.MachineSpecDisk)
		d.Key = disk.Key
		d.Volume = disk.Volume
		if pool.Spec.Type == "rbd" {
			d.Type = "network"
			d.Device = "disk"

			d.Network = new(spec.MachineSpecDiskNetwork)
			d.Network.Protocol = pool.Spec.Type
			d.Network.Hosts = pool.Spec.Source.Hosts
			d.Network.Auth = *pool.Spec.Source.Auth

			disks = append(disks, d)
			rockferry.MachineEnsureUniqueDiskTargets(disks, rockferry.MachineDiskTargetBaseVD)
			continue
		}

		if pool.Spec.Type == "dir" {
			d.Type = "file"
			d.Device = "disk"

			d.File = new(spec.MachineSpecDiskFile)

			disks = append(disks, d)
			rockferry.MachineEnsureUniqueDiskTargets(disks, rockferry.MachineDiskTargetBaseSD)
			continue
		}
	}

	if t.Request.Spec.Cdrom.Volume != "" {
		volume, err := executor.Rockferry.StorageVolumes().Get(ctx, t.Request.Spec.Cdrom.Volume, nil)
		if err != nil {
			return nil, err
		}

		cdrom := new(spec.MachineSpecDisk)

		cdrom.Key = volume.Spec.Key
		cdrom.Volume = volume.Id

		// This could probably be more clean
		cdrom.File = new(spec.MachineSpecDiskFile)
		cdrom.Device = "cdrom"
		cdrom.Type = "file"

		disks = append(disks, cdrom)
		rockferry.MachineEnsureUniqueDiskTargets(disks, rockferry.MachineDiskTargetBaseSD)
	}

	return disks, nil
}

func (t *CreateVirtualMachineTask) createNetworkInterfaces(ctx context.Context, executor *Executor) ([]*spec.MachineSpecInterface, error) {
	interfaces := make([]*spec.MachineSpecInterface, 1)

	networks, err := executor.Rockferry.Networks().List(ctx, t.Request.Spec.Network, nil)
	if err != nil {
		return nil, err
	}

	network := networks[0]

	mac, err := mac.Generate()
	if err != nil {
		return nil, err
	}

	interfaces[0] = new(spec.MachineSpecInterface)
	interfaces[0].Mac = mac
	interfaces[0].Model = "virtio"

	interfaces[0].Network = new(string)
	*interfaces[0].Network = network.Spec.Name

	interfaces[0].Bridge = new(string)
	*interfaces[0].Bridge = network.Spec.Bridge.Name

	return interfaces, nil
}

func (t *CreateVirtualMachineTask) downloadPossibleResources(machineSpec *spec.MachineSpec) error {
	basePath := fmt.Sprintf("/var/rockferry/assets")

	if err := ensurePath(basePath); err != nil {
		return err
	}

	for k, v := range t.Request.Annotations {
		if k == "kernel.download" {
			// /var/rockferry/assets/${UUID}
			path := fmt.Sprintf("%s/%s", basePath, uuid.NewString())
			if err := downloadFile(path, v); err != nil {
				return err
			}

			machineSpec.Boot.Kernel = pointer.To(path)
		}

		if k == "initramfs.download" {
			// /var/rockferry/assets/${UUID}
			path := fmt.Sprintf("%s/%s", basePath, uuid.NewString())
			if err := downloadFile(path, v); err != nil {
				return err
			}

			machineSpec.Boot.Initramfs = pointer.To(path)
		}

		if k == "kernel.cmdline" {
			machineSpec.Boot.Cmdline = pointer.To(v)
		}
	}

	return nil
}

func (t *CreateVirtualMachineTask) Execute(ctx context.Context, executor *Executor) error {
	// NOTE: Used to annotate storage volumes with the vm id. This is useful for deletion.
	vmId := uuid.NewString()

	disks, err := t.createVmDisks(ctx, executor)
	if err != nil {
		return err
	}

	interfaces, err := t.createNetworkInterfaces(ctx, executor)
	if err != nil {
		return err
	}

	machineSpec := new(spec.MachineSpec)

	if err := t.downloadPossibleResources(machineSpec); err != nil {
		return err
	}

	machineSpec.Boot.Order = []string{"hd", "cdrom"}

	machineSpec.Name = t.Request.Spec.Name
	machineSpec.Topology = t.Request.Spec.Topology
	machineSpec.Disks = disks
	machineSpec.Interfaces = interfaces

	res := new(rockferry.Machine)

	res.Annotations = map[string]string{}
	res.Annotations["machinerequest.id"] = t.Request.Id

	res.Id = vmId
	res.Kind = rockferry.ResourceKindMachine
	res.Owner = new(rockferry.OwnerRef)
	// TODO: Do not hardcode this
	res.Owner.Id = executor.NodeId
	res.Owner.Kind = rockferry.ResourceKindNode

	res.Status.State = spec.MachineStatusStateBooting

	res.Spec = *machineSpec

	if err := executor.Libvirt.CreateDomain(vmId, machineSpec); err != nil {
		return err
	}

	return executor.Rockferry.Machines().Create(ctx, res)
}

func (t *CreateVirtualMachineTask) Resource() *rockferry.Resource[any, any] {
	return t.Request.Generic()
}

type DeleteVmTask struct {
	Machine *rockferry.Machine
}

func (t *DeleteVmTask) Execute(ctx context.Context, e *Executor) error {
	if t.Machine.Status.State == spec.MachineStatusStateRunning || t.Machine.Status.State == spec.MachineStatusStateBooting {
		if err := e.Libvirt.DestroyDomain(t.Machine.Id); err != nil {
			return err
		}
	}

	if err := e.Libvirt.UndefineDomain(t.Machine.Id); err != nil {
		return err
	}

	// Cleanup, yay
	for _, disk := range t.Machine.Spec.Disks {
		if disk.Volume == "" {
			continue
		}

		if err := e.Rockferry.StorageVolumes().Delete(ctx, disk.Volume); err != nil {
			fmt.Println("failed to delete storage volume", err)
			continue
		}
	}

	return e.Rockferry.MachineRequests().Delete(ctx, t.Machine.Annotations["machinerequest.id"])
}

func (t *DeleteVmTask) Repeats() *time.Duration {
	return nil
}

type SyncMachineStatusesTask struct {
}

func (t *SyncMachineStatusesTask) Execute(ctx context.Context, e *Executor) error {
	iface := e.Rockferry.Machines()

	owner := new(rockferry.OwnerRef)
	owner.Id = e.NodeId
	owner.Kind = rockferry.ResourceKindNode
	machines, err := iface.List(ctx, "", owner)
	if err != nil && err != rockferry.ErrorNotFound {
		return err
	}

	for _, machine := range machines {
		// TODO: If a machine exists in rockferry but not libvirt we are out of sync.
		// 		 We need logic to sync machines as well.
		if !e.Libvirt.DomainExists(machine.Id) {
			continue
		}

		status, err := e.Libvirt.SyncDomainStatus(machine.Id)
		if err != nil {
			fmt.Println("failed to sync machine status", err)
			continue
		}

		copy := new(rockferry.Machine)
		*copy = *machine

		copy.Status = *status

		if err := iface.Patch(ctx, machine, copy); err != nil {
			fmt.Println("failed to patch machine", err)
			continue
		}
	}

	return nil
}

func (t *SyncMachineStatusesTask) Repeats() *time.Duration {
	timeout := time.Second * 2
	return &timeout
}
