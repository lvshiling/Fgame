package player

import (
	"fgame/fgame/game/global"
	zhenfatemplate "fgame/fgame/game/zhenfa/template"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
)

func (m *PlayerZhenFaDataManager) GmSetZhenFaLevel(zhenFaType zhenfatypes.ZhenFaType, level int32) {
	zhenFaTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaTempalte(zhenFaType, level)
	if zhenFaTemplate == nil {
		return
	}
	obj := m.GetZhenFaByType(zhenFaType)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.level = level
	obj.updateTime = now
	obj.SetModified()
}

func (m *PlayerZhenFaDataManager) GmSetZhenQiLevel(zhenFaType zhenfatypes.ZhenFaType, zhenFaPos zhenfatypes.ZhenQiType, level int32) {
	zhenQiTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaZhenQiTemplate(zhenFaType, zhenFaPos, level)
	if zhenQiTemplate == nil {
		return
	}
	obj := m.GetZhenQi(zhenFaType, zhenFaPos)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.number = level
	obj.updateTime = now
	obj.SetModified()
}

func (m *PlayerZhenFaDataManager) GmSetZhenQiXianHuoLevel(zhenFaType zhenfatypes.ZhenFaType, level int32) {
	zhenFaXianHuoTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaXianHuoTemplate(zhenFaType, level)
	if zhenFaXianHuoTemplate == nil {
		return
	}
	obj := m.GetZhenQiXianHuo(zhenFaType)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	obj.level = level
	obj.updateTime = now
	obj.SetModified()
}
