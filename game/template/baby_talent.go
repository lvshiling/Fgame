package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	babytypes "fgame/fgame/game/baby/types"
	"fmt"
)

//宝宝天赋配置
type BabyTalentTemplate struct {
	*BabyTalentTemplateVO
	nextTemp  *BabyTalentTemplate
	skillType babytypes.SkillType
}

func (t *BabyTalentTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyTalentTemplate) GetNextTemplate() *BabyTalentTemplate {
	return t.nextTemp
}

func (t *BabyTalentTemplate) GetSkillType() babytypes.SkillType {
	return t.skillType
}

func (t *BabyTalentTemplate) PatchAfterCheck() {

}

func (t *BabyTalentTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//下一阶强化
	if t.NextId != 0 {
		if t.NextId-t.Id != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}

		to := template.GetTemplateService().Get(t.NextId, (*BabyTalentTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}
		t.nextTemp = to.(*BabyTalentTemplate)
	}

	//技能类型
	t.skillType = babytypes.SkillType(t.Type)

	return nil
}

func (t *BabyTalentTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if !t.skillType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//权值
	err = validator.MinValidate(float64(t.Rate), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Rate)
		return template.NewTemplateFieldError("Rate", err)
	}

	return nil
}

func (t *BabyTalentTemplate) FileName() string {
	return "tb_baobao_tianfu.json"
}

func init() {
	template.Register((*BabyTalentTemplate)(nil))
}
