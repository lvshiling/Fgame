package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type MarryXinWuSuitTemplate struct {
	*MarryXinWuSuitTemplateVO
}

func (t *MarryXinWuSuitTemplate) TemplateId() int {
	return t.Id
}

func (t *MarryXinWuSuitTemplate) FileName() string {
	return "tb_marry_xinwu_suit.json"
}

func (t *MarryXinWuSuitTemplate) Check() (err error) {
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}
	err = validator.MinValidate(float64(t.Num), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Num)
		err = template.NewTemplateFieldError("Num", err)
		return
	}

	if t.Type != 2 {
		err = fmt.Errorf("[%d] Type", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	return
}

func (t *MarryXinWuSuitTemplate) Patch() (err error) {
	return nil
}

func (t *MarryXinWuSuitTemplate) PatchAfterCheck() {
	return
}

func init() {
	template.Register((*MarryXinWuSuitTemplate)(nil))
}
