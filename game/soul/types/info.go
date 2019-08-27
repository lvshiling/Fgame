package types

type SoulInfo struct {
	SoulTag        int32 `json:"soulTag"`
	Level          int32 `json:"level"`
	Experience     int32 `json:"experience"`
	AwakenOrder    int32 `json:"awakenOrder"`
	IsAwaken       int32 `json:"isAwaken"`
	StrengthenLevl int32 `json:"strengthenLevel"`
	StrengthenPro  int32 `json:"strengthenPro"`
}

type AllSoulInfo struct {
	SoulList    []*SoulInfo `json:"soulList"`
	EmbedIdList []int32     `json:"embedIdList"`
}
