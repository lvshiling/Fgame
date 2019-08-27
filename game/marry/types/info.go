package types

type MarryInfo struct {
	SpouseId   int64  `json:spouseId`
	SpouseName string `json:spouseName`
	Ring       int32  `json:"ring"`
	RLevel     int32  `json:"rLevel"`
	RNum       int32  `json:"rNum"`
	RProgress  int32  `json:"rProgress"`
	TLevel     int32  `json:"tLevel"`
	TNum       int32  `json:"tNum"`
	TProgress  int32  `json:"tProgress"`
	IsProposal int32  `json:"isProposal"`
	Status     int32  `json:"status"`
	DLevel     int32  `json:"dLevel"`
	DExp       int32  `json:"dExp"`
	MarryCount int32  `json:"marryCount"`
}

type MarryHeroism struct {
	name    string
	heroism int32
}
