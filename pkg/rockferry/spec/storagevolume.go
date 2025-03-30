package spec

type StorageVolumeType string

const (
	StorageVolumeTypeDiskImage = "image"
	StorageVolumeTypeIso       = "iso"
)

type StorageVolumeSpec struct {
	Name       string            `json:"name"`
	Type       StorageVolumeType `json:"type"`
	Capacity   uint64            `json:"capacity"`
	Allocation uint64            `json:"allocation"`
	Key        string            `json:"key"`
	Pool       string            `json:"pool"`
}
