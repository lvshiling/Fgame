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
	skilltypes "fgame/fgame/game/skill/types"
	playerwing "fgame/fgame/game/wing/player"
	"fgame/fgame/game/wing/wing"
	playerxianzuncard "fgame/fgame/game/xianzuncard/player"
	xianzuncardtemplate "fgame/fgame/game/xianzuncard/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeWing, WingPropertyEffect)
}

//战翼作用器
func WingPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeWing) {
		return
	}
	xianZunManager := p.GetPlayerDataManager(types.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	wingManager := p.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)

	wingInfo := wingManager.GetWingInfo()
	wingTrialInfo := wingManager.GetWingTrialInfo()
	advancedId := wingInfo.AdvanceId
	//战翼试用属性
	if wingTrialInfo.TrialOrderId != 0 {
		advancedId = 1
	}
	wingTemplate := wing.GetWingService().GetWingNumber(int32(advancedId))
	//战翼系统默认不开启
	if wingTemplate == nil {
		return
	}

	//幻化丹食丹等级
	unrealLevel := wingInfo.UnrealLevel
	huanHuaTemplate := wing.GetWingService().GetWingHuanHuaTemplate(unrealLevel)
	if huanHuaTemplate != nil {

		hp := int64(huanHuaTemplate.Hp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, hp)
		attack := int64(huanHuaTemplate.Attack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, attack)
		defence := int64(huanHuaTemplate.Defence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, defence)
	}

	// 计算加成的万分比

	xianZunMap := xianZunManager.GetXianZunCardObjectMap()
	for typ, obj := range xianZunMap {
		if !obj.IsActivite() {
			continue
		}

		xianZunTemp := xianzuncardtemplate.GetXianZunCardTemplateService().GetXianZunCardTemplate(typ)
		if xianZunTemp == nil {
			continue
		}

		hp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP) + int64(xianZunTemp.WingAttrAddPercent)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, hp)
		attack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack) + int64(xianZunTemp.WingAttrAddPercent)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, attack)
		defence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend) + int64(xianZunTemp.WingAttrAddPercent)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, defence)
	}

	//战翼属性
	if wingTemplate.GetBattleAttrTemplate() != nil {
		for typ, val := range wingTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

	//非进阶战翼
	wingOtherMap := wingManager.GetWingOtherMap()
	for _, wingTypeOtherMap := range wingOtherMap {
		for wingOtherId, wo := range wingTypeOtherMap {
			wingTemplate := wing.GetWingService().GetWing(int(wingOtherId))
			//非进阶战翼属性
			if wingTemplate.GetBattleAttrTemplate() == nil {
				continue
			}

			skinHp := int64(0)
			skinAttack := int64(0)
			skinDefence := int64(0)
			if wingTemplate.WingUpstarBeginId != 0 && wo.Level != 0 {
				wingUpstarTemplate := wingTemplate.GetWingUpstarByLevel(wo.Level)
				skinHp = int64(wingUpstarTemplate.Hp)
				skinAttack = int64(wingUpstarTemplate.Attack)
				skinDefence = int64(wingUpstarTemplate.Defence)

				//战翼基础全属性万分比
				if wingUpstarTemplate.WingPercent != 0 {
					oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(wingUpstarTemplate.WingPercent)+oldHp)
					oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(wingUpstarTemplate.WingPercent)+oldAttack)
					oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(wingUpstarTemplate.WingPercent)+oldDefence)

				}
			}

			for typ, val := range wingTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
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
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeWingStone, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeWingStone, prop)
}
