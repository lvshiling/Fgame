package effect

import (
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeLingyuAura, DoushenLingyuPropertyEffect)
}

//斗神领域技能光环作用器
func DoushenLingyuPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if p.GetAllianceId() == 0 {
		return
	}

	al := alliance.GetAllianceService().GetAlliance(p.GetAllianceId())
	if al == nil {
		return
	}

	levelLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDouShenLevelLimit)
	if p.GetLevel() < levelLimit {
		return
	}

	for _, mem := range al.GetDouShenList() {
		lingyuTemp := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(mem.GetLingyuId())
		if lingyuTemp == nil {
			continue
		}
		skillTemplate := lingyuTemp.GetSkillTemplate()
		if skillTemplate == nil {
			continue
		}
		//领域属性
		for typ, val := range skillTemplate.GetAttrAuraTemplate().GetAllBattleProperty() {
			total := prop.GetGlobal(typ)
			total += val
			prop.SetGlobal(typ, total)
		}
	}
}
