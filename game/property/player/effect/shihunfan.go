package effect

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	skilltypes "fgame/fgame/game/skill/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeShiHunFan, ShiHunFanPropertyEffect)
}

//噬魂幡作用器
func ShiHunFanPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeShiHunFan) {
		return
	}

	manager := p.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shihunfanInfo := manager.GetShiHunFanInfo()
	advancedId := shihunfanInfo.AdvanceId
	shihunfanTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(int32(advancedId))
	//噬魂幡系统默认不开启 advancedId=0
	if shihunfanTemplate == nil {
		return
	}
	hp := int64(0)
	attack := int64(0)
	defence := int64(0)

	//噬魂幡食丹等级
	danLevel := shihunfanInfo.DanLevel
	danTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanDan(danLevel)
	if danTemplate != nil {
		danHp := int64(danTemplate.Hp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, danHp)
		danAttack := int64(danTemplate.Attack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, danAttack)
		danDefence := int64(danTemplate.Defence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, danDefence)
	}

	//噬魂幡属性
	hp += int64(shihunfanTemplate.Hp)
	attack += int64(shihunfanTemplate.Attack)
	defence += int64(shihunfanTemplate.Defence)

	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, hp)
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, attack)
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, defence)

	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeShiHunFan, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeShiHunFanAdvancedSkill, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeShiHunFanSystemSkill, prop)
}
