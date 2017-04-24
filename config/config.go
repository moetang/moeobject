package config

type Config struct {
	MachId  uint16         `json:"mach_id"`
	Volumes []VolumeConfig `json:"volumes"`
}

type VolumeConfig struct {
	VolId uint16 `json:"vol_id"`
	Path  string `json:"path"`
}
