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
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotemplate "fgame/fgame/game/tianmo/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeTianMoTi, TianMoPropertyEffect)
}

//天魔体作用器
func TianMoPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeTianMo) {
		return
	}
	tianmoManager := p.GetPlayerDataManager(types.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianmoInfo := tianmoManager.GetTianMoInfo()
	advancedId := tianmoInfo.AdvanceId
	tianmoTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(advancedId)
	if tianmoTemplate == nil {
		return
	}

	hp := int64(0)
	attack := int64(0)
	defence := int64(0)
	//培养丹食丹等级
	culLevel := tianmoInfo.TianMoDanLevel
	danTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoDan(culLevel)
	if danTemplate != nil {
		hp += int64(danTemplate.Hp)
		attack += int64(danTemplate.Attack)
		defence += int64(danTemplate.Defence)

		hp += prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, hp)
		attack += prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, attack)
		defence += prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, defence)
	}

	//天魔体属性
	for typ, val := range tianmoTemplate.GetBattleAttrMap() {
		total := prop.GetBase(typ)
		total += val
		prop.SetBase(typ, total)
	}
	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeTianMoTi, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeTianMoAdvancedSkill, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeTianMoSystemSkill, prop)
}
