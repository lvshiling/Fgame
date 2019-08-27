package common

import (
	"fgame/fgame/game/lingtongdev/template"
	"fgame/fgame/game/lingtongdev/types"
)

//皮肤信息
type LingTongDevSkinInfo struct {
	SeqId int32 `json:"seqId"`
	Level int32 `json:"level"`
	UpPro int32 `json:"upPro"`
}

//灵童养成类信息
type LingTongDevInfo struct {
	ClassType     int32                  `json:"classType"`
	AdvanceId     int                    `json:"advanceId"`
	SeqId         int32                  `json:"seqId"`
	UnrealLevel   int32                  `json:"unrealLevel"`
	UnrealPro     int32                  `json:"unrealPro"`
	CulLevel      int32                  `json:"culLevel"`
	CulPro        int32                  `json:"culPro"`
	TongLingLevel int32                  `json:"tongLingLevel"`
	TongLingPro   int32                  `json:"tongLingPro"`
	SkinList      []*LingTongDevSkinInfo `json:"skinList"`
}

func (m *LingTongDevInfo) GetSeqId() int32 {
	if m.SeqId == 0 {
		classType := types.LingTongDevSysType(m.ClassType)
		lingTongDevTemplate := template.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, m.AdvanceId)
		if lingTongDevTemplate != nil {
			return int32(lingTongDevTemplate.TemplateId())
		}
	}
	return m.SeqId
}

type AllLingTongDevInfo struct {
	LingTongDevList []*LingTongDevInfo `json:"lingTongDevList"`
}
