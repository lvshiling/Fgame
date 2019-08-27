package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeLingTongFashion, LingTongFashionPropertyEffect)
}

//灵童时装作用器
func LingTongFashionPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeLingTongHuanHua) {
		return
	}
	manager := p.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	activateMap := manager.GetActivateFashionMap()
	fashionTrialObj := manager.GetFashionTrialObject()

	if fashionTrialObj != nil && !fashionTrialObj.GetIsExpire() {
		fashionId := fashionTrialObj.GetTrialFashionId()
		lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
		if lingTongFashionTemplate != nil {
			battlePropertyMap := lingTongFashionTemplate.GetBattleProperty()
			for typ, val := range battlePropertyMap {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}
	}

	//获取关联模块
	propertyManager := p.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	equipmentBaseModule := propertyManager.GetModule(playerpropertytypes.PlayerPropertyEffectorTypeLingTong)
	//获取装备
	fashionPropertySegment := equipmentBaseModule.GetExternalPropertySegment(playerpropertytypes.PlayerPropertyEffectorTypeLingTongFashion)
	fashionPropertySegment.Clear()
	for fashionId, wo := range activateMap {
		lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
		if lingTongFashionTemplate != nil {
			battlePropertyMap := lingTongFashionTemplate.GetBattleProperty()
			for typ, val := range battlePropertyMap {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}

		if lingTongFashionTemplate.LingTongUpstarId != 0 && wo.GetUpgradeLevel() != 0 {
			upstarTemplate := lingTongFashionTemplate.GetLingTongFashionUpstarByLevel(wo.GetUpgradeLevel())
			//基础全属性万分比
			upstarPercent := int64(upstarTemplate.Percent)
			if upstarPercent != 0 {
				//装备基础全属性万分比
				oldHp := fashionPropertySegment.Get(uint(propertytypes.BattlePropertyTypeMaxHP))
				fashionPropertySegment.Set(uint(propertytypes.BattlePropertyTypeMaxHP), upstarPercent+oldHp)
				oldAttack := fashionPropertySegment.Get(uint(propertytypes.BattlePropertyTypeAttack))
				fashionPropertySegment.Set(uint(propertytypes.BattlePropertyTypeAttack), upstarPercent+oldAttack)
				oldDefence := fashionPropertySegment.Get(uint(propertytypes.BattlePropertyTypeDefend))
				fashionPropertySegment.Set(uint(propertytypes.BattlePropertyTypeDefend), upstarPercent+oldDefence)
			}

			battlePropertyMap := upstarTemplate.GetBattlePropertyMap()
			for typ, val := range battlePropertyMap {
				value := val + prop.GetGlobal(typ)
				prop.SetGlobal(typ, value)
			}
		}
	}
}
