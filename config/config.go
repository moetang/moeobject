package config

type Config struct {
	MachId  uint16         `json:"mach_id"`
	PhyFileConfigs []PhyFileConfig `json:"phyfiles"`
}

type PhyFileConfig struct {
	VolId uint16 `json:"phyfile_id"`
	Path  string `json:"path"`
}
