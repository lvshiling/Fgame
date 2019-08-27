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
	playerxianti "fgame/fgame/game/xianti/player"
	"fgame/fgame/game/xianti/xianti"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeXianTi, XianTiPropertyEffect)
}

//仙体作用器
func XianTiPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeXianTi) {
		return
	}
	xianTiManager := p.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xianTiInfo := xianTiManager.GetXianTiInfo()
	advancedId := xianTiInfo.AdvanceId
	xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(advancedId))
	if xianTiTemplate == nil {
		return
	}

	for typ, val := range xianTiTemplate.GetBattleProperty() {
		total := prop.GetBase(typ)
		total += val
		prop.SetBase(typ, total)
	}

	//幻化丹食丹等级
	unrealLevel := xianTiInfo.UnrealLevel
	huanHuaTemplate := xianti.GetXianTiService().GetXianTiHuanHuaTemplate(unrealLevel)
	if huanHuaTemplate != nil {
		hp := int64(huanHuaTemplate.Hp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, hp)
		attack := int64(huanHuaTemplate.Attack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, attack)
		defence := int64(huanHuaTemplate.Defence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, defence)
	}

	//非进阶仙体
	xianTiOtherMap := xianTiManager.GetXianTiOtherMap()
	for _, xianTiTypeOtherMap := range xianTiOtherMap {
		for xianTiOtherId, wo := range xianTiTypeOtherMap {
			xianTiTemplate := xianti.GetXianTiService().GetXianTi(int(xianTiOtherId))

			if xianTiTemplate.XianTiUpstarBeginId != 0 && wo.Level != 0 {
				xianTiUpstarTemplate := xianTiTemplate.GetXianTiUpstarByLevel(wo.Level)
				skinHp := int64(xianTiUpstarTemplate.Hp)
				skinAttack := int64(xianTiUpstarTemplate.Attack)
				skinDefence := int64(xianTiUpstarTemplate.Defence)

				//仙体基础全属性万分比
				if xianTiUpstarTemplate.XianTiPercent != 0 {
					oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(xianTiUpstarTemplate.XianTiPercent)+oldHp)
					oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(xianTiUpstarTemplate.XianTiPercent)+oldAttack)
					oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(xianTiUpstarTemplate.XianTiPercent)+oldDefence)
				}

				hp := int64(skinHp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
				prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, hp)
				attack := int64(skinAttack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
				prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, attack)
				defence := int64(skinDefence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
				prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, defence)
			}

			for typ, val := range xianTiTemplate.GetBattleProperty() {
				total := prop.GetGlobal(typ)
				total += val
				prop.SetGlobal(typ, total)
			}
		}
	}

	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeXianTi, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeXianTi, prop)
}
