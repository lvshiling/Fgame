package effect

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeLingTongMount, LingTongMountPropertyEffect)
}

//灵骑作用器
func LingTongMountPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeLingTongMount) {
		return
	}
	classType := lingtongdevtypes.LingTongDevSysTypeLingQi
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

	for typ, val := range lingTongDevTemplate.GetBattlePropertyMap() {
		total := prop.GetBase(typ)
		total += val
		prop.SetBase(typ, total)
	}

	//幻化丹食丹等级
	unrealLevel := lingTongDevInfo.GetUnrealLevel()
	lingTongDevHuanHuaTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevHuanHuaTemplate(classType, unrealLevel)
	if lingTongDevHuanHuaTemplate != nil {
		battlePropertyMap := lingTongDevHuanHuaTemplate.GetBattlePropertyMap()
		for typ, val := range battlePropertyMap {
			value := val + prop.GetGlobal(typ)
			prop.SetGlobal(typ, value)
		}
	}

	//培养丹食丹等级
	culLevel := lingTongDevInfo.GetCulLevel()
	lingTongDevPeiYangTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevPeiYangTemplate(classType, culLevel)
	if lingTongDevPeiYangTemplate != nil {
		battlePropertyMap := lingTongDevPeiYangTemplate.GetBattlePropertyMap()
		for typ, val := range battlePropertyMap {
			value := val + prop.GetGlobal(typ)
			prop.SetGlobal(typ, value)
		}
	}

	//通灵
	tonglingLevel := lingTongDevInfo.GetTongLingLevel()
	lingTongDevTongLingTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTongLingTemplate(classType, tonglingLevel)
	if lingTongDevTongLingTemplate != nil {
		percent := lingTongDevTongLingTemplate.GetTongLingPercent()
		oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(percent)+oldHp)
		oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(percent)+oldAttack)
		oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(percent)+oldDefence)
	}

	//非进阶
	lingTongDevContainer := manager.GetLingTongDevOtherMap(classType)
	if lingTongDevContainer != nil {
		for _, lingTongDevTypeOtherMap := range lingTongDevContainer.GetOtherMap() {
			for seqId, wo := range lingTongDevTypeOtherMap {
				lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, int(seqId))

				if lingTongDevTemplate.GetUpstarBeginId() != 0 && wo.GetLevel() != 0 {
					upstarTemplate := lingTongDevTemplate.GetLingTongDevUpstarByLevel(wo.GetLevel())

					//基础全属性万分比
					upstarPercent := int64(upstarTemplate.GetUpstarPercent())
					if upstarPercent != 0 {
						oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
						prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, upstarPercent+oldHp)
						oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
						prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, upstarPercent+oldAttack)
						oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
						prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, upstarPercent+oldDefence)
					}

					battlePropertyMap := upstarTemplate.GetBattlePropertyMap()
					for typ, val := range battlePropertyMap {
						value := val + prop.GetGlobal(typ)
						prop.SetGlobal(typ, value)
					}
				}

				for typ, val := range lingTongDevTemplate.GetBattlePropertyMap() {
					total := prop.GetGlobal(typ)
					total += val
					prop.SetGlobal(typ, total)
				}
			}
		}
	}

	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeLingTongMount, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeLingTongMountSystemSkill, prop)
}
