package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//兑换钱
type MoneyTemplate struct {
	*MoneyTemplateVO
}

func (t *MoneyTemplate) TemplateId() int {
	return t.Id
}

func (t *MoneyTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *MoneyTemplate) PatchAfterCheck() {

}

func (t *MoneyTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//验证 Money
	err = validator.MinValidate(float64(t.Money), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Money)
		err = template.NewTemplateFieldError("Money", err)
		return
	}
	//验证 Money
	err = validator.MinValidate(float64(t.Gold), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Gold)
		err = template.NewTemplateFieldError("Gold", err)
		return
	}
	return nil
}

func (t *MoneyTemplate) FileName() string {
	return "tb_money.json"
}

func init() {
	template.Register((*MoneyTemplate)(nil))
}
