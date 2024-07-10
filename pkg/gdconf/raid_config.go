package gdconf

type RaidConfig struct {
	RaidID           uint32   `json:"RaidID"`
	HardLevel        uint32   `json:"HardLevel"`
	RaidTagList      []string `json:"RaidTagList"`
	UnlockWorldLevel []uint32 `json:"UnlockWorldLevel"`
}
