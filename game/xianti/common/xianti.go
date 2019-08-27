package common

import "fgame/fgame/game/xianti/xianti"

//皮肤信息
type XianTiSkinInfo struct {
	XianTiId int32 `json:"xianTiId"`
	Level    int32 `json:"level"`
	UpPro    int32 `json:"upPro"`
}

//仙体信息
type XianTiInfo struct {
	AdvanceId   int               `json:"advanceId"`
	XianTiId    int32             `json:"xianTiId"`
	UnrealLevel int32             `json:"unrealLevel"`
	UnrealPro   int32             `json:"unrealPro"`
	SkinList    []*XianTiSkinInfo `json:"skinList"`
}

func (m *XianTiInfo) GetXianTiId() int32 {
	if m.XianTiId == 0 {
		xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(m.AdvanceId))
		if xianTiTemplate != nil {
			return int32(xianTiTemplate.TemplateId())
		}
	}
	return m.XianTiId
}
