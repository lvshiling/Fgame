package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type SystemAwakeLevelTemplate struct {
	*SystemAwakeLevelTemplateVO
	nextTemplate *SystemAwakeLevelTemplate
}

func (t *SystemAwakeLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *SystemAwakeLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return
}

func (t *SystemAwakeLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*SystemAwakeLevelTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemplate, _ = to.(*SystemAwakeLevelTemplate)

		diff := t.nextTemplate.Level - int32(t.Level)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	// 验证等级
	err = validator.MinValidate(float64(t.Level), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	// 验证银两
	err = validator.MinValidate(float64(t.UseSilver), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseSilver)
		err = template.NewTemplateFieldError("UseSilver", err)
		return
	}

	//验证need_number
	err = validator.MinValidate(float64(t.UseItemCount), 1, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItemCount)
		err = template.NewTemplateFieldError("UseItemCount", err)
		return
	}

	//验证 hp
	err = validator.MinValidate(float64(t.Hp), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(t.Attack), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(t.Defence), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(t.TimesMin), float64(0), true, float64(t.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(t.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(t.AddMin), float64(0), true, float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(t.AddMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	return
}

func (t *SystemAwakeLevelTemplate) PatchAfterCheck() {

}

func (t *SystemAwakeLevelTemplate) FileName() string {
	return "tb_system_juexing_level.json"
}

func init() {
	template.Register((*SystemAwakeLevelTemplate)(nil))
}
