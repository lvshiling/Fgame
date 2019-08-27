package effect

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	skilltypes "fgame/fgame/game/skill/types"
)

// func init() {
// 	playerpropertytypes.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeShenfa, ShenfaPropertyEffect)
// }

// //身法作用器
// func ShenfaPropertyEffect(p player.Player, prop *propertycommon.BattlePropertySegment) {
// 	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeShenfa) {
// 		return
// 	}
// 	shenfaManager := p.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
// 	shenfaInfo := shenfaManager.GetShenfaInfo()
// 	advancedId := shenfaInfo.AdvanceId
// 	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(int32(advancedId))
// 	//身法系统默认不开启 advancedId=0
// 	if shenfaTemplate == nil {
// 		return
// 	}
// 	//其它身法皮肤属性
// 	for _, shenfaId := range shenfaInfo.UnrealList {
// 		unrealShenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(shenfaId)
// 		if unrealShenfaTemplate == nil {
// 			continue
// 		}
// 		if unrealShenfaTemplate.GetTyp() == shenfatypes.ShenfaTypeAdvanced {
// 			continue
// 		}

// 		for typ, val := range unrealShenfaTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
// 			total := prop.Get(typ)
// 			total += val
// 			prop.Set(typ, total)
// 		}
// 	}

// 	hp := int64(0)
// 	attack := int64(0)
// 	defence := int64(0)
// 	//幻化丹食丹等级
// 	unrealLevel := shenfaInfo.UnrealLevel
// 	huanHuaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaHuanHuaTemplate(unrealLevel)
// 	if huanHuaTemplate != nil {
// 		hp += int64(huanHuaTemplate.Hp)
// 		attack += int64(huanHuaTemplate.Attack)
// 		defence += int64(huanHuaTemplate.Defence)
// 	}

// 	//身法属性
// 	if shenfaTemplate.GetBattleAttrTemplate() != nil {
// 		for typ, val := range shenfaTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
// 			switch typ {
// 			case propertytypes.BattlePropertyTypeMaxHP:
// 				{
// 					val += hp
// 					break
// 				}
// 			case propertytypes.BattlePropertyTypeAttack:
// 				{
// 					val += attack
// 					break
// 				}
// 			case propertytypes.BattlePropertyTypeDefend:
// 				{
// 					val += defence
// 					break
// 				}
// 			}

// 			total := prop.Get(typ)
// 			total += val
// 			prop.Set(typ, total)
// 		}
// 	}

// 	//非进阶身法
// 	shenfaOtherMap := shenfaManager.GetShenfaOtherMap()
// 	for _, shenfaObj := range shenfaOtherMap {
// 		for _, shenfaOtherId := range shenfaObj.ShenfaList {
// 			shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenfaOtherId))
// 			//非进阶身法属性
// 			if shenfaTemplate.GetBattleAttrTemplate() == nil {
// 				continue
// 			}
// 			for typ, val := range shenfaTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
// 				total := prop.Get(typ)
// 				total += val
// 				prop.Set(typ, total)
// 			}
// 		}
// 	}

// }

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeShenfa, ShenfaPropertyEffect)
}

//身法作用器
func ShenfaPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeShenfa) {
		return
	}
	shenfaManager := p.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenfaInfo := shenfaManager.GetShenfaInfo()
	advancedId := shenfaInfo.AdvanceId
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(int32(advancedId))
	//身法系统默认不开启 advancedId=0
	if shenfaTemplate == nil {
		return
	}

	//幻化丹食丹等级
	unrealLevel := shenfaInfo.UnrealLevel
	huanHuaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaHuanHuaTemplate(unrealLevel)
	if huanHuaTemplate != nil {
		hp := int64(huanHuaTemplate.Hp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, hp)
		attack := int64(huanHuaTemplate.Attack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, attack)
		defence := int64(huanHuaTemplate.Defence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, defence)
	}

	//身法属性
	if shenfaTemplate.GetBattleAttrTemplate() != nil {
		for typ, val := range shenfaTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

	//非进阶领域
	shenfaOtherMap := shenfaManager.GetShenfaOtherMap()
	for _, shenFaTypeOtherMap := range shenfaOtherMap {
		for shenFaOtherId, wo := range shenFaTypeOtherMap {
			shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenFaOtherId))
			//非进阶领域属性
			if shenfaTemplate.GetBattleAttrTemplate() == nil {
				continue
			}

			skinHp := int64(0)
			skinAttack := int64(0)
			skinDefence := int64(0)
			if shenfaTemplate.ShenfaUpstarBeginId != 0 && wo.Level != 0 {
				shenFaUpstarTemplate := shenfaTemplate.GetShenFaUpstarByLevel(wo.Level)
				skinHp = int64(shenFaUpstarTemplate.Hp)
				skinAttack = int64(shenFaUpstarTemplate.Attack)
				skinDefence = int64(shenFaUpstarTemplate.Defence)

				//身法基础全属性万分比
				if shenFaUpstarTemplate.ShenFaPercent != 0 {
					oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(shenFaUpstarTemplate.ShenFaPercent)+oldHp)
					oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(shenFaUpstarTemplate.ShenFaPercent)+oldAttack)
					oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(shenFaUpstarTemplate.ShenFaPercent)+oldDefence)
				}
			}

			for typ, val := range shenfaTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
				switch typ {
				case propertytypes.BattlePropertyTypeMaxHP:
					{
						val += skinHp
						break
					}
				case propertytypes.BattlePropertyTypeAttack:
					{
						val += skinAttack
						break
					}
				case propertytypes.BattlePropertyTypeDefend:
					{
						val += skinDefence
						break
					}
				}
				total := prop.GetGlobal(typ)
				total += val
				prop.SetGlobal(typ, total)
			}
		}
	}
	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeShenFa, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeShenfa, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeShenFaSystemSkill, prop)
}
