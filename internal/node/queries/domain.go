package queries

import (
	"encoding/xml"
	"fmt"
	"net"
	"strconv"

	"github.com/digitalocean/go-libvirt"
	"github.com/eskpil/rockferry/pkg/rockferry/spec"
	"github.com/eskpil/rockferry/pkg/virtwrap/domain"
	"github.com/google/uuid"
)

// TODO: There are a lot more configuration options here which can be set.
func (c *Client) CreateDomain(id string, spec *spec.MachineSpec) error {
	schema := new(domain.Schema)

	schema.Name = spec.Name
	schema.Type = "kvm"

	schema.UUID = id

	schema.Memory.Unit = "bytes"
	schema.Memory.Value = spec.Topology.Memory

	schema.VCPU = new(domain.VCPU)
	schema.VCPU.CPUs = uint32(spec.Topology.Cores) * uint32(spec.Topology.Threads)
	schema.VCPU.Placement = "static"

	schema.CPU.Topology = new(domain.CPUTopology)
	schema.CPU.Topology.Cores = uint32(spec.Topology.Cores)
	schema.CPU.Topology.Threads = uint32(spec.Topology.Threads)
	schema.CPU.Topology.Sockets = 1
	schema.CPU.Mode = "host-passthrough"

	schema.Features = new(domain.Features)
	schema.Features.ACPI = new(domain.FeatureEnabled)
	schema.Features.APIC = new(domain.FeatureEnabled)

	schema.Devices.Emulator = "/usr/bin/qemu-system-x86_64"

	schema.OS.Type.Arch = "x86_64"
	schema.OS.Type.Machine = "pc-q35-7.2"
	schema.OS.Type.OS = "hvm"

	for _, e := range spec.Boot.Order {
		schema.OS.BootOrder = append(schema.OS.BootOrder, domain.Boot{Dev: e})
	}

	if spec.Boot.Kernel != nil && spec.Boot.Initramfs != nil {
		schema.OS.Initrd = *spec.Boot.Initramfs
		schema.OS.Kernel = *spec.Boot.Kernel
		if spec.Boot.Cmdline != nil {
			schema.OS.KernelArgs = *spec.Boot.Cmdline
		}
	}

	schema.SysInfo = new(domain.SysInfo)
	schema.SysInfo.Type = "smbios"
	schema.SysInfo.System = append(schema.SysInfo.System, domain.Entry{Name: "manufacturer", Value: "rockferry"})
	schema.SysInfo.System = append(schema.SysInfo.System, domain.Entry{Name: "product", Value: "rockferry"})
	schema.SysInfo.System = append(schema.SysInfo.System, domain.Entry{Name: "version", Value: "alpha-1"})
	schema.SysInfo.System = append(schema.SysInfo.System, domain.Entry{Name: "uuid", Value: id})

	for _, d := range spec.Disks {
		disk := new(domain.Disk)

		if d.Type == "network" {
			disk.Type = "network"
			disk.Device = d.Device

			disk.Driver = new(domain.DiskDriver)
			disk.Driver.Name = "qemu"
			disk.Driver.Type = "raw"

			disk.Auth = new(domain.DiskAuth)

			disk.Auth.Username = d.Network.Auth.Username
			disk.Auth.Secret = new(domain.DiskSecret)
			disk.Auth.Secret.Type = d.Network.Auth.Type
			disk.Auth.Secret.UUID = d.Network.Auth.Secret

			disk.Source.Protocol = d.Network.Protocol
			disk.Source.Name = d.Key
			disk.Source.Host = new(domain.DiskSourceHost)
			disk.Source.Host.Name = d.Network.Hosts[0].Name
			disk.Source.Host.Port = d.Network.Hosts[0].Port

			disk.Target.Bus = "virtio"
			// TODO: Create a function which returns unique device names
			disk.Target.Device = "vda"
		}

		if d.Type == "file" {
			disk.Type = "file"
			disk.Device = d.Device

			disk.Source.File = d.Key

			disk.Driver = new(domain.DiskDriver)
			disk.Driver.Name = "qemu"
			disk.Driver.Type = "raw"

			disk.Target.Bus = "sata"
			disk.Target.Device = "sda"

		}

		schema.Devices.Disks = append(schema.Devices.Disks, *disk)
	}

	for _, i := range spec.Interfaces {
		iface := new(domain.Interface)

		iface.MAC = new(domain.MAC)
		iface.MAC.MAC = i.Mac
		iface.Type = "network"
		iface.Source.Network = *i.Network
		iface.Model = new(domain.Model)
		iface.Model.Type = "virtio"

		schema.Devices.Interfaces = append(schema.Devices.Interfaces, *iface)
	}

	qga := new(domain.Channel)

	qga.Type = "unix"
	qga.Source = new(domain.ChannelSource)
	qga.Source.Mode = "bind"
	qga.Source.Path = fmt.Sprintf("/var/lib/libvirt/qemu/%s.agent", id)

	qga.Target = new(domain.ChannelTarget)
	qga.Target.Type = "virtio"
	qga.Target.Name = "org.qemu.guest_agent.0"

	vnc := new(domain.Graphics)

	schema.Devices.Channels = append(schema.Devices.Channels, *qga)

	vnc.Type = "vnc"
	//vnc.AutoPort = "yes"
	vnc.Websocket = "-1"
	vnc.Listen = new(domain.GraphicsListen)

	vnc.Listen.Type = "address"
	vnc.Listen.Address = "0.0.0.0"

	schema.Devices.Graphics = append(schema.Devices.Graphics, *vnc)

	bytes, err := xml.Marshal(schema)
	if err != nil {
		return err
	}

	dom, err := c.v.DomainDefineXML(string(bytes))
	if err != nil {
		return err
	}

	return c.v.DomainCreate(dom)
}

func (c *Client) DestroyDomain(id string) error {
	domId := uuid.MustParse(id)

	domain, err := c.v.DomainLookupByUUID(libvirt.UUID(domId))
	if err != nil {
		return err
	}

	return c.v.DomainDestroy(domain)
}

func (c *Client) DomainExists(id string) bool {
	domId := uuid.MustParse(id)

	_, err := c.v.DomainLookupByUUID(libvirt.UUID(domId))
	if err != nil {
		return false
	}

	return true
}

func (c *Client) GetDomainState(id string) (spec.MachineStatusState, error) {
	domId := uuid.MustParse(id)

	dom, err := c.v.DomainLookupByUUID(libvirt.UUID(domId))
	if err != nil {
		return "", err
	}

	state, _, err := c.v.DomainGetState(dom, 0)
	switch state {
	case 1:
		{
			return spec.MachineStatusStateRunning, nil
		}
	case 6:
		{
			return spec.MachineStatusStateCrashed, nil
		}
	default:
		{
			return spec.MachineStatusStateStopped, nil
		}
	}
}

func (c *Client) SyncDomainStatus(id string) (*spec.MachineStatus, error) {
	status := new(spec.MachineStatus)

	domId := uuid.MustParse(id)

	dom, err := c.v.DomainLookupByUUID(libvirt.UUID(domId))
	if err != nil {
		return nil, err
	}

	state, _, err := c.v.DomainGetState(dom, 0)
	switch state {
	case 1:
		{
			status.State = spec.MachineStatusStateRunning
			break
		}
	case 6:
		{
			status.State = spec.MachineStatusStateCrashed
			break
		}
	default:
		{
			status.State = spec.MachineStatusStateStopped
			break
		}
	}

	// TODO: Best to avoid this completely. DomainGetXMLDesc is a very heavy operation.
	xmlSchema, err := c.v.DomainGetXMLDesc(dom, 0)
	if err != nil {
		return nil, err
	}

	schema := new(domain.Schema)
	if err := xml.Unmarshal([]byte(xmlSchema), schema); err != nil {
		return nil, err
	}

	for _, graphicDevice := range schema.Devices.Graphics {
		if graphicDevice.Type != "vnc" {
			continue
		}

		// TODO: When both are present create a seperate entry for both.
		if graphicDevice.Websocket != "" {
			port, err := strconv.ParseInt(graphicDevice.Websocket, 10, 16)
			if err != nil {
				return nil, err
			}

			status.VNC = append(status.VNC, spec.MachineStatusVNC{Port: int32(port), Type: spec.MachineStatusVNCTypeWebsocket})
		} else {
			status.VNC = append(status.VNC, spec.MachineStatusVNC{Port: graphicDevice.Port, Type: spec.MachineStatusVNCTypeNative})
		}
	}

	interfaces, err := c.v.DomainInterfaceAddresses(dom, 1, 0)
	if err != nil {
		status.Errors = append(status.Errors, err.Error())
		return status, nil
	}

	status.Interfaces = make([]spec.MachineStatusInterface, len(interfaces))

	for i, iface := range interfaces {
		addrs := []spec.MachineStatusIp{}

		for _, a := range iface.Addrs {
			ip := net.ParseIP(a.Addr)

			addr := spec.MachineStatusIp{
				Ip:      ip.String(),
				Private: ip.IsPrivate(),
			}

			if addr.Private {
				status.ReachableIps = append(status.ReachableIps, addr)
			}

			addrs = append(addrs, addr)
		}

		status.Interfaces[i] = spec.MachineStatusInterface{
			Name:  iface.Name,
			Mac:   iface.Hwaddr[0],
			Addrs: addrs,
		}
	}

	return status, nil
}

func (c *Client) StartDomain(id string) error {
	domId := uuid.MustParse(id)

	dom, err := c.v.DomainLookupByUUID(libvirt.UUID(domId))
	if err != nil {
		return err
	}

	return c.v.DomainCreate(dom)
}

func (c *Client) ShutdownDomain(id string) error {
	domId := uuid.MustParse(id)

	dom, err := c.v.DomainLookupByUUID(libvirt.UUID(domId))
	if err != nil {
		return err
	}

	return c.v.DomainShutdownFlags(dom, 4)
}

func (c *Client) UndefineDomain(id string) error {
	domId := uuid.MustParse(id)

	dom, err := c.v.DomainLookupByUUID(libvirt.UUID(domId))
	if err != nil {
		return err
	}

	return c.v.DomainUndefine(dom)
}
