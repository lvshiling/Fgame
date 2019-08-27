package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

// [
//   {
//     "id": 1,
//     "next_id": 2,
//     "type": 1,
//     "level": 1,
//     "cost_yinliang": 300,
//     "cost_yuanbao": 0,
//     "cost_item_id": "0",
//     "cost_item_count": "0",
//     "force": 100,
//     "damage_value_base": 5,
//     "damage_value": 0,
//     "spell_damage": 10000,
//     "spell_power": 5000
//   },
// ]

type SkillLevelTemplate struct {
	//模板数据
	*SkillLevelTemplateVO
	nextSkillLevelTemplate *SkillLevelTemplate
	//damageAttack
	damageAttack float64
}

func (slt *SkillLevelTemplate) GetNextSkillLevelTemplate() *SkillLevelTemplate {
	return slt.nextSkillLevelTemplate
}

func (slt *SkillLevelTemplate) GetDamageAttack() float64 {
	return slt.damageAttack
}

func (slt *SkillLevelTemplate) TemplateId() int {
	return slt.Id
}
func (slt *SkillLevelTemplate) PatchAfterCheck() {
}
func (slt *SkillLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(slt.FileName(), slt.TemplateId(), err)
			return
		}
	}()
	slt.damageAttack = float64(slt.SpellDamage) / common.MAX_RATE * float64(slt.SpellPower) / common.MAX_RATE

	//nextId
	if slt.NextId != 0 {
		tempNextSkillLevelTemplate := template.GetTemplateService().Get(int(slt.NextId), (*SkillLevelTemplate)(nil))
		if tempNextSkillLevelTemplate == nil {
			err = fmt.Errorf("[%d] invalid", slt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		slt.nextSkillLevelTemplate = tempNextSkillLevelTemplate.(*SkillLevelTemplate)
		if slt.nextSkillLevelTemplate.Typ != slt.Typ {
			err = fmt.Errorf("templateId [%d] of [%d] invalid", slt.nextSkillLevelTemplate.TemplateId(), slt.nextSkillLevelTemplate.Typ)
			return template.NewTemplateFieldError("type", err)
		}
		diffLevel := slt.nextSkillLevelTemplate.Level - slt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("templateId [%d] of [%d] invalid", slt.nextSkillLevelTemplate.TemplateId(), slt.nextSkillLevelTemplate.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}
	return nil
}

func (slt *SkillLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(slt.FileName(), slt.TemplateId(), err)
			return
		}
	}()
	//消耗银两
	if err = validator.MinValidate(float64(slt.CostSilver), float64(0), true); err != nil {
		return
	}
	//消耗元宝
	if err = validator.MinValidate(float64(slt.CostGold), float64(0), true); err != nil {
		return
	}
	//添加战斗力
	if err = validator.MinValidate(float64(slt.Force), float64(0), true); err != nil {
		return
	}
	//伤害
	if err = validator.MinValidate(float64(slt.DamageValueBase), float64(0), true); err != nil {
		return
	}
	//伤害成长
	if err = validator.MinValidate(float64(slt.DamageValue), float64(0), true); err != nil {
		return
	}
	//技能伤害
	if err = validator.MinValidate(float64(slt.SpellDamage), float64(0), true); err != nil {
		return
	}
	//技能伤害
	if err = validator.MinValidate(float64(slt.SpellPower), float64(0), true); err != nil {
		return
	}

	return nil
}
func (slt *SkillLevelTemplate) FileName() string {
	return "tb_skill_level.json"
}

func init() {
	template.Register((*SkillLevelTemplate)(nil))
}
