package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//转生礼包折扣配置
type CircleBargainTemplate struct {
	*CircleBargainTemplateVO
}

func (t *CircleBargainTemplate) TemplateId() int {
	return t.Id
}

func (t *CircleBargainTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *CircleBargainTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 Type
	err = validator.MinValidate(float64(t.Type), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 ItemCountMin
	err = validator.MinValidate(float64(t.ItemCountMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ItemCountMin)
		err = template.NewTemplateFieldError("ItemCountMin", err)
		return
	}

	//验证 ItemCountMax
	err = validator.MinValidate(float64(t.ItemCountMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ItemCountMax)
		err = template.NewTemplateFieldError("ItemCountMax", err)
		return
	}

	//验证 Discount
	err = validator.MinValidate(float64(t.Discount), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Discount)
		err = template.NewTemplateFieldError("Discount", err)
		return
	}

	return nil
}

func (t *CircleBargainTemplate) PatchAfterCheck() {

}

func (t *CircleBargainTemplate) FileName() string {
	return "tb_circle_bargain.json"
}

func init() {
	template.Register((*CircleBargainTemplate)(nil))
}
