package effect

import (
	"fgame/fgame/game/player"
	propertycommon "fgame/fgame/game/property/common"
	propertytypes "fgame/fgame/game/property/types"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
)

//模块技能作用器
func skillByModulePropertyEffect(p player.Player, sft skilltypes.SkillFirstType, prop *propertycommon.SystemPropertySegment) {

	passiveSkillMap := p.GetSkills(skilltypes.SkillSecondTypePassive)
	for _, passiveSkill := range passiveSkillMap {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(passiveSkill.GetSkillId(), passiveSkill.GetLevel())
		if skillTemplate.GetSkillFirstType() != sft {
			continue
		}
		attrTemplate := skillTemplate.GetAttrTemplate()
		for typ, val := range attrTemplate.GetAllBattleProperty() {
			total := prop.GetGlobal(typ)
			total += val
			prop.SetGlobal(typ, total)
		}
	}

	//添加战力
	// skillManager := p.GetPlayerDataManager(playertypes.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	for _, ski := range p.GetAllSkills() {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(ski.GetSkillId())
		if skillTemplate.GetSkillFirstType() != sft {
			continue
		}
		if skillTemplate.IsStatic() {
			skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(ski.GetSkillId(), ski.GetLevel())
			total := prop.GetGlobal(propertytypes.BattlePropertyTypeForce)
			total += int64(skillTemplate.AddPower)
			prop.SetGlobal(propertytypes.BattlePropertyTypeForce, total)
		} else {
			skillLevelTemplate := skillTemplate.GetSkillByLevel(ski.GetLevel())
			total := prop.GetGlobal(propertytypes.BattlePropertyTypeForce)
			total += int64(skillLevelTemplate.Force)
			prop.SetGlobal(propertytypes.BattlePropertyTypeForce, total)
		}
	}

	for _, ski := range p.GetAllSkills() {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(ski.GetSkillId())
		if skillTemplate.GetSkillFirstType() != sft {
			continue
		}
		if skillTemplate.GetAttrTemplate() == nil {
			continue
		}
		if skillTemplate.IsStatic() {
			skillTemplate = skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(ski.GetSkillId(), ski.GetLevel())
		}
		battlePropertyPercentMap := skillTemplate.GetAttrTemplate().GetBattlePropertyPercent()
		for typ, percent := range battlePropertyPercentMap {
			oldVal := prop.GetGlobalPercent(typ)
			prop.SetGlobalPercent(typ, int64(oldVal+percent))
		}
	}
}
