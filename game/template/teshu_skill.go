package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

type TeShuSkillTemplate struct {
	*TeShuSkillTemplateVO
	skillTemplate *SkillTemplate
}

func (t *TeShuSkillTemplate) TemplateId() int {
	return t.Id
}

func (t *TeShuSkillTemplate) GetSkillTemplate() *SkillTemplate {
	return t.skillTemplate
}

func (t *TeShuSkillTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *TeShuSkillTemplate) PatchAfterCheck() {

}

func (t *TeShuSkillTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//技能类型
	if err = validator.MinValidate(float64(t.SkillType), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("SkillType", err)
		return
	}

	//技能Id
	if err = validator.MinValidate(float64(t.SkillId), float64(0), false); err != nil {
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}
	//关联技能模板
	temp := template.GetTemplateService().Get(int(t.SkillId), (*SkillTemplate)(nil))
	skillTemp, _ := temp.(*SkillTemplate)
	if skillTemp == nil {
		err = fmt.Errorf("TeShuSkillTemplate[%d] invalid", t.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}

	if skillTemp.GetSkillFirstType() != skilltypes.SkillFirstTypeCastingSoul {
		err = fmt.Errorf("TeShuSkillTemplate[%d] invalid", t.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}
	t.skillTemplate = skillTemp

	return nil
}

func (t *TeShuSkillTemplate) FileName() string {
	return "tb_teshu_skill.json"
}

func init() {
	template.Register((*TeShuSkillTemplate)(nil))
}
