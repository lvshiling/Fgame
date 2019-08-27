package types

//皮肤信息
type ShenfaSkinInfo struct {
	ShenfaId int32 `json:"shenfaId"`
	Level    int32 `json:"level"`
	UpPro    int32 `json:"upPro"`
}

type ShenfaInfo struct {
	AdvanceId   int32             `json:"advanceId"`
	ShenfaId    int32             `json:"shenfaId"`
	UnrealLevel int32             `json:"unrealLevel"`
	UnrealPro   int32             `json:"unrealPro"`
	SkinList    []*ShenfaSkinInfo `json:"skinList"`
}
