package common

import "fgame/fgame/game/fabao/template"

//皮肤信息
type FaBaoSkinInfo struct {
	FaBaoId int32 `json:"faBaoId"`
	Level   int32 `json:"level"`
	UpPro   int32 `json:"upPro"`
}

type FaBaoInfo struct {
	AdvanceId     int32            `json:"advanceId"`
	FaBaoId       int32            `json:"faBaoId"`
	UnrealLevel   int32            `json:"unrealLevel"`
	UnrealPro     int32            `json:"unrealPro"`
	TongLingLevel int32            `json:"tonglingLevel"`
	TongLingPro   int32            `json:"tonglingPro"`
	SkinList      []*FaBaoSkinInfo `json:"skinList"`
}

func (w *FaBaoInfo) GetFaBaoId() int32 {
	if w.FaBaoId == 0 {
		faBaoTemplate := template.GetFaBaoTemplateService().GetFaBaoNumber(w.AdvanceId)
		if faBaoTemplate != nil {
			return int32(faBaoTemplate.TemplateId())
		}
	}
	return w.FaBaoId
}
