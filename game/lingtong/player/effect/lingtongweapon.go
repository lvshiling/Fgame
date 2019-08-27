package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	lingtongplayerproperty "fgame/fgame/game/lingtong/player/property"
	lingtongplayertypes "fgame/fgame/game/lingtong/player/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
)

func init() {
	lingtongplayerproperty.RegisterLingTongPropertyEffector(lingtongplayertypes.LingTongPropertyEffectorTypeWeapon, WeaponPropertyEffect)
}

//初始作用器
func WeaponPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeLingTongWeapon) {
		return
	}
	classType := lingtongdevtypes.LingTongDevSysTypeLingBing
	manager := p.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	if lingTongDevInfo == nil {
		return
	}
	advancedId := lingTongDevInfo.GetAdvancedId()
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, advancedId)
	//系统默认不开启
	if lingTongDevTemplate == nil {
		return
	}

	for typ, val := range lingTongDevTemplate.GetLingTongBattlePropertyMap() {
		total := prop.GetBase(typ)
		total += val
		prop.SetBase(typ, total)
	}

	//幻化丹食丹等级
	unrealLevel := lingTongDevInfo.GetUnrealLevel()
	lingTongDevHuanHuaTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevHuanHuaTemplate(classType, unrealLevel)
	if lingTongDevHuanHuaTemplate != nil {
		battlePropertyMap := lingTongDevHuanHuaTemplate.GetLingTongBattlePropertyMap()
		for typ, val := range battlePropertyMap {
			value := val + prop.GetGlobal(typ)
			prop.SetGlobal(typ, value)
		}
	}

	//培养丹食丹等级
	culLevel := lingTongDevInfo.GetCulLevel()
	lingTongDevPeiYangTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevPeiYangTemplate(classType, culLevel)
	if lingTongDevPeiYangTemplate != nil {
		battlePropertyMap := lingTongDevPeiYangTemplate.GetLingTongBattlePropertyMap()
		for typ, val := range battlePropertyMap {
			value := val + prop.GetGlobal(typ)
			prop.SetGlobal(typ, value)
		}
	}

	//非进阶
	lingTongDevContainer := manager.GetLingTongDevOtherMap(classType)
	if lingTongDevContainer != nil {
		for _, lingTongDevTypeOtherMap := range lingTongDevContainer.GetOtherMap() {
			for seqId, wo := range lingTongDevTypeOtherMap {
				lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, int(seqId))

				if lingTongDevTemplate.GetUpstarBeginId() != 0 && wo.GetLevel() != 0 {
					upstarTemplate := lingTongDevTemplate.GetLingTongDevUpstarByLevel(wo.GetLevel())

					battlePropertyMap := upstarTemplate.GetLingTongBattlePropertyMap()
					for typ, val := range battlePropertyMap {
						value := val + prop.GetGlobal(typ)
						prop.SetGlobal(typ, value)
					}
				}

				for typ, val := range lingTongDevTemplate.GetLingTongBattlePropertyMap() {
					total := prop.GetGlobal(typ)
					total += val
					prop.SetGlobal(typ, total)
				}
			}
		}
	}

}
