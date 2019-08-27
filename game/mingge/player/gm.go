package player

import (
	"fgame/fgame/game/global"
	minggetemplate "fgame/fgame/game/mingge/template"
	minggetypes "fgame/fgame/game/mingge/types"
)

func (m *PlayerMingGeDataManager) GmSetMingPanLevel(mingGeSubType minggetypes.MingGeAllSubType, number int32, star int32) {
	mingPanTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingPanTemplate(mingGeSubType, number, star)
	if mingPanTemplate == nil {
		return
	}
	obj := m.GetMingGePanRefinedByType(mingGeSubType)
	if obj == nil {
		obj = m.initMingPanRefined(mingGeSubType)
	}
	now := global.GetGame().GetTimeService().Now()
	obj.number = number
	obj.star = star
	obj.refinedNum = 0
	obj.refinedPro = 0
	obj.updateTime = now
	obj.SetModified()
}

func (m *PlayerMingGeDataManager) GmSetMingLi(mingGongType minggetypes.MingGongType, mingGongSubType minggetypes.MingGongAllSubType, slotType minggetypes.MingLiSlotType, propertyType minggetypes.MingGePropertyType) (mingGongTypeMap map[minggetypes.MingGongType]bool) {
	obj := m.GetMingGeMingLiByTypeAndSubType(mingGongType, mingGongSubType)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	mingLiMap := obj.GetMingLiMap()
	mingLiInfo := mingLiMap[slotType]
	mingLiInfo.MingGeProperty = propertyType
	obj.updateTime = now
	obj.SetModified()

	//判断命宫是否激活
	level := m.p.GetLevel()
	zhuanShu := m.p.GetZhuanSheng()
	mingGongTypeMap = m.CheckMingGongActivate(level, zhuanShu)
	return
}

func (m *PlayerMingGeDataManager) GmSetMingLiNum(mingGongType minggetypes.MingGongType, mingGongSubType minggetypes.MingGongAllSubType, slotType minggetypes.MingLiSlotType, num int32) (mingGongTypeMap map[minggetypes.MingGongType]bool) {
	obj := m.GetMingGeMingLiByTypeAndSubType(mingGongType, mingGongSubType)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	mingLiMap := obj.GetMingLiMap()
	mingLiInfo := mingLiMap[slotType]
	mingLiInfo.Times = num
	obj.updateTime = now
	obj.SetModified()

	//判断命宫是否激活
	level := m.p.GetLevel()
	zhuanShu := m.p.GetZhuanSheng()
	mingGongTypeMap = m.CheckMingGongActivate(level, zhuanShu)
	return
}
