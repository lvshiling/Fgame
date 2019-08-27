package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
)

//玩具套装属性配置
type BabyToySuitTemplate struct {
	*BabyToySuitTemplateVO
}

func (t *BabyToySuitTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyToySuitTemplate) Patch() (err error) {

	return nil
}

func (t *BabyToySuitTemplate) PatchAfterCheck() {

}

func (t *BabyToySuitTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//套装数量
	if err = validator.MinValidate(float64(t.Num), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("Num", err)
		return
	}

	//属性
	if err = validator.MinValidate(float64(t.Value1), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("Value1", err)
		return
	}

	return nil
}

func (t *BabyToySuitTemplate) FileName() string {
	return "tb_baobao_wanju_suit.json"
}

func init() {
	template.Register((*BabyToySuitTemplate)(nil))
}
