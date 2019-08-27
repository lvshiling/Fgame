package types

//皮肤信息
type LingyuSkinInfo struct {
	LingyuId int32 `json:"lingyuId"`
	Level    int32 `json:"level"`
	UpPro    int32 `json:"upPro"`
}

type LingyuInfo struct {
	AdvanceId   int32             `json:"advanceId"`
	LingyuId    int32             `json:"lingyuId"`
	UnrealLevel int32             `json:"unrealLevel"`
	UnrealPro   int32             `json:"unrealPro"`
	SkinList    []*LingyuSkinInfo `json:"skinList"`
}
