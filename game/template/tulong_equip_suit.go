package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
)

//屠龙套装属性配置
type TuLongEquipSuitTemplate struct {
	*TuLongEquipSuitTemplateVO
}

func (t *TuLongEquipSuitTemplate) TemplateId() int {
	return t.Id
}

func (t *TuLongEquipSuitTemplate) Patch() (err error) {

	return nil
}

func (t *TuLongEquipSuitTemplate) PatchAfterCheck() {

}

func (t *TuLongEquipSuitTemplate) Check() (err error) {
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

func (t *TuLongEquipSuitTemplate) FileName() string {
	return "tb_tulongequip_suit.json"
}

func init() {
	template.Register((*TuLongEquipSuitTemplate)(nil))
}
