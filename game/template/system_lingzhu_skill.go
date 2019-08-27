package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type SystemLingZhuSkillTemplate struct {
	*SystemLingZhuSkillTemplateVO
	nextSystemLingZhuSkillTemplate *SystemLingZhuSkillTemplate
}

func (t *SystemLingZhuSkillTemplate) TemplateId() int {
	return t.Id
}

//检查有效性
func (t *SystemLingZhuSkillTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//灵童ID
	err = validator.MinValidate(float64(t.LingtongId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingtongId)
		return template.NewTemplateFieldError("LingtongId", err)
	}

	//等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//所需灵珠等级
	err = validator.MinValidate(float64(t.NeedLingzhuLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedLingzhuLevel)
		return template.NewTemplateFieldError("NeedLingzhuLevel", err)
	}

	//下一等级
	err = validator.MinValidate(float64(t.NextId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NextId)
		return template.NewTemplateFieldError("NextId", err)
	}

	//检查nextId可不可靠
	if t.nextSystemLingZhuSkillTemplate != nil {
		diff := t.nextSystemLingZhuSkillTemplate.Level - t.Level
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	return
}

//组合成需要的数据
func (t *SystemLingZhuSkillTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	if t.NextId == 0 {
		t.nextSystemLingZhuSkillTemplate = nil
	} else {
		temp := template.GetTemplateService().Get(int(t.NextId), (*SystemLingZhuSkillTemplate)(nil))
		t.nextSystemLingZhuSkillTemplate, _ = temp.(*SystemLingZhuSkillTemplate)
		if t.nextSystemLingZhuSkillTemplate == nil {
			err = fmt.Errorf("SystemLingZhuSkillTemplate[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("MagicConditionParameter", err)
		}
	}

	return
}

//检验后组合
func (t *SystemLingZhuSkillTemplate) PatchAfterCheck() {
}

func (t *SystemLingZhuSkillTemplate) FileName() string {
	return "tb_lingtong_lingzhu_skill.json"
}

func init() {
	template.Register((*SystemLingZhuSkillTemplate)(nil))
}
