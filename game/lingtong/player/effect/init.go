package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongplayerproperty "fgame/fgame/game/lingtong/player/property"
	lingtongplayertypes "fgame/fgame/game/lingtong/player/types"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
)

func init() {
	lingtongplayerproperty.RegisterLingTongPropertyEffector(lingtongplayertypes.LingTongPropertyEffectorTypeInit, InitPropertyEffect)
}

//初始作用器
func InitPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeLingTong) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := manager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		return
	}
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongObj.GetLingTongId())
	if lingTongTemplate == nil {
		return
	}
	battlePropertyMap := lingTongTemplate.GetLingTongBattlePropertyMap()
	for k, v := range battlePropertyMap {
		prop.SetBase(k, v)
	}
	lingTongInfObj, _ := manager.GetLingTongInfo(lingTongObj.GetLingTongId())
	//培养丹食丹等级
	culLevel := lingTongInfObj.GetPeiYangLevel()
	lingTongPeiYangTemplate := lingTongTemplate.GetLingTongPeiYangByLevel(culLevel)
	if lingTongPeiYangTemplate != nil {
		battlePropertyMap := lingTongPeiYangTemplate.GetLingTongBattlePropertyMap()
		for typ, val := range battlePropertyMap {
			value := val + prop.GetGlobal(typ)
			prop.SetGlobal(typ, value)
		}
	}

	//升级
	level := lingTongInfObj.GetLevel()
	lingTongShengJiTemplate := lingTongTemplate.GetLingTongShengJiByLevel(level)
	if lingTongShengJiTemplate != nil {
		battlePropertyMap := lingTongShengJiTemplate.GetLingTongBattlePropertyMap()
		for typ, val := range battlePropertyMap {
			value := val + prop.GetBase(typ)
			prop.SetBase(typ, value)
		}
	}
}
