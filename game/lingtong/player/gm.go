package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
)

//仅gm使用 灵童激活
func (m *PlayerLingTongDataManager) GmLingTongActivate(lingTongId int32) (flag bool) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	_, flag = m.GetLingTongInfo(lingTongId)
	if flag {
		return
	}
	m.LingTongActivate(lingTongId)
	flag = true
	return
}

//仅gm使用 灵童升级
func (m *PlayerLingTongDataManager) GmLingTongShengJi(lingTongId int32, level int32) (flag bool) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	lingTongShengJiTemplate := lingTongTemplate.GetLingTongShengJiByLevel(level)
	if lingTongShengJiTemplate == nil {
		return
	}

	lingTongInfo, flag := m.GetLingTongInfo(lingTongId)
	if !flag {
		return
	}
	diffLevel := level - lingTongInfo.upgradeLevel
	now := global.GetGame().GetTimeService().Now()
	lingTongInfo.upgradeLevel = level
	lingTongInfo.upgradeNum = 0
	lingTongInfo.upgradePro = 0
	lingTongInfo.updateTime = now
	lingTongInfo.SetModified()
	flag = true

	//总等级
	m.lingTongObj.level += diffLevel
	m.lingTongObj.updateTime = now
	m.lingTongObj.SetModified()
	gameevent.Emit(lingtongeventtypes.EventTypeLingTongLevelChanged, m.p, m.lingTongObj)
	return
}

//仅gm使用 灵童培养
func (m *PlayerLingTongDataManager) GmLingTongPeiYang(lingTongId int32, level int32) (flag bool) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return
	}
	lingTongPeiYangTemplate := lingTongTemplate.GetLingTongPeiYangByLevel(level)
	if lingTongPeiYangTemplate == nil {
		return
	}

	lingTongInfo, flag := m.GetLingTongInfo(lingTongId)
	if !flag {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	lingTongInfo.peiYangLevel = level
	lingTongInfo.peiYangNum = 0
	lingTongInfo.peiYangPro = 0
	lingTongInfo.updateTime = now
	lingTongInfo.SetModified()
	flag = true
	return
}

//仅gm使用 灵童时装激活
func (m *PlayerLingTongDataManager) GmLingTongFashionActivate(fashionId int32) (flag bool) {
	lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(int32(fashionId))
	//时装灵童
	if lingTongFashionTemplate == nil {
		return
	}
	flag = lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if flag {
		return
	}
	fashionObj := m.GetFashionInfoById(fashionId)
	if fashionObj != nil {
		return
	}

	_, flag = m.FashionActive(fashionId)
	return
}

//仅gm使用 灵童时装升星
func (m *PlayerLingTongDataManager) GmLingTongFashionUpstar(fashionId int32, level int32) (flag bool) {
	lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(int32(fashionId))
	//时装灵童
	if lingTongFashionTemplate == nil {
		return
	}
	lingTongFashionUpstarTemplate := lingTongFashionTemplate.GetLingTongFashionUpstarByLevel(level)
	if lingTongFashionUpstarTemplate == nil {
		return
	}
	flag = lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if flag {
		return
	}
	fashionInfo := m.GetFashionInfoById(fashionId)
	if fashionInfo == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	fashionInfo.upgradeLevel = level
	fashionInfo.upgradeNum = 0
	fashionInfo.upgradePro = 0
	fashionInfo.updateTime = now
	fashionInfo.SetModified()
	flag = true
	return
}
