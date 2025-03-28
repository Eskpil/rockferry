package spec

type MachineStatusState string
type MachineStatusVNCType string

const (
	MachineStatusStateRunning   MachineStatusState = "running"
	MachineStatusStateRebooting                    = "rebooting"
	MachineStatusStateCrashed                      = "crashed"
	MachineStatusStateStopped                      = "stopped"
	MachineStatusStateShutdown                     = "shutdown"
	MachineStatusStateBooting                      = "booting"

	MachineStatusVNCTypeWebsocket MachineStatusVNCType = "websocket"
	MachineStatusVNCTypeNative                         = "native"
)

type MachineSpecInterface struct {
	Mac   string `json:"mac"`
	Model string `json:"model"`

	Network *string `json:"network"`
	Bridge  *string `json:"bridge"`
}

type MachineSpecDiskTarget struct {
	Dev string `json:"dev"`
}

type MachineSpecDiskFile struct {
}

type MachineSpecDiskNetwork struct {
	Hosts []*StoragePoolSpecSourceHost `json:"hosts"`
	Auth  StoragePoolSpecSourceAuth    `json:"auth"`

	Protocol string `json:"type"`
}

type MachineSpecDisk struct {
	Device string                 `json:"device"`
	Type   string                 `json:"type"`
	Key    string                 `json:"key"`
	Volume string                 `json:"volume"`
	Target *MachineSpecDiskTarget `json:"target"`

	File    *MachineSpecDiskFile    `json:"file,omitempty"`
	Network *MachineSpecDiskNetwork `json:"network,omitempty"`
}

type MachineSpecBoot struct {
	Order []string `json:"order"`

	Kernel    *string `json:"kernel,omitempty"`
	Initramfs *string `json:"initramfs,omitempty"`
	Cmdline   *string `json:"cmdline,omitempty"`
}

type MachineSpec struct {
	Name     string   `json:"name"`
	Topology Topology `json:"topology"`

	Boot MachineSpecBoot `json:"boot"`

	Disks      []*MachineSpecDisk      `json:"disks"`
	Interfaces []*MachineSpecInterface `json:"interfaces"`
}

type MachineStatusVNC struct {
	Port int32                `json:"port"`
	Type MachineStatusVNCType `json:"type"`
}

type MachineStatusIp struct {
	Ip      string `json:"ip"`
	Private bool   `json:"private"`
}

type MachineStatusInterface struct {
	Name  string            `json:"name"`
	Mac   string            `json:"mac"`
	Addrs []MachineStatusIp `json:"addrs"`
}

type MachineStatus struct {
	State  MachineStatusState `json:"state"`
	Errors []string           `json:"errors"`

	VNC        []MachineStatusVNC       `json:"vnc"`
	Interfaces []MachineStatusInterface `json:"interfaces"`

	ReachableIps []MachineStatusIp `json:"reachable_ips"`
}
