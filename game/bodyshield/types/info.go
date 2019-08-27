package types

//护体盾信息
type BodyShieldInfo struct {
	AdvancedId     int32 `json:"advancedId"`
	JinjiadanLevel int32 `json:"jinjiadanLevel"`
	JinjiadanPro   int32 `json:"jinjiadanPro"`
}

type ShieldInfo struct {
	ShieldId int32 `json:"shieldId"`
	Progress int32 `json:"progress"`
}
