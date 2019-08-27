package common

import "fgame/fgame/game/wing/wing"

//皮肤信息
type WingSkinInfo struct {
	WingId int32 `json:"wingId"`
	Level  int32 `json:"level"`
	UpPro  int32 `json:"upPro"`
}

type WingInfo struct {
	AdvanceId   int32           `json:"advanceId"`
	WingId      int32           `json:"wingId"`
	UnrealLevel int32           `json:"unrealLevel"`
	UnrealPro   int32           `json:"unrealPro"`
	SkinList    []*WingSkinInfo `json:"skinList"`
}

func (w *WingInfo) GetWingId() int32 {
	if w.WingId == 0 {
		wingTemplate := wing.GetWingService().GetWingNumber(w.AdvanceId)
		if wingTemplate != nil {
			return int32(wingTemplate.TemplateId())
		}
	}
	return w.WingId
}

type FeatherInfo struct {
	FeatherId int32 `json:"featherId"`
	Progress  int32 `json:"progress"`
}
