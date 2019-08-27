package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
)

//套装属性配置
type GoldEquipSuitTemplate struct {
	*GoldEquipSuitTemplateVO
}

func (t *GoldEquipSuitTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldEquipSuitTemplate) Patch() (err error) {

	return nil
}

func (t *GoldEquipSuitTemplate) PatchAfterCheck() {

}

func (t *GoldEquipSuitTemplate) Check() (err error) {
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

func (t *GoldEquipSuitTemplate) FileName() string {
	return "tb_goldequip_suit.json"
}

func init() {
	template.Register((*GoldEquipSuitTemplate)(nil))
}
