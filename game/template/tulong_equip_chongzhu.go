package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
)

//屠龙重铸配置
type TuLongEquipRongHeTemplate struct {
	*TuLongEquipRongHeTemplateVO
}

func (t *TuLongEquipRongHeTemplate) TemplateId() int {
	return t.Id
}

func (t *TuLongEquipRongHeTemplate) Patch() (err error) {

	return nil
}

func (t *TuLongEquipRongHeTemplate) PatchAfterCheck() {

}

func (t *TuLongEquipRongHeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//转生数
	if err = validator.MinValidate(float64(t.Level), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	return nil
}

func (t *TuLongEquipRongHeTemplate) FileName() string {
	return "tb_tulongequip_ronghe.json"
}

func init() {
	template.Register((*TuLongEquipRongHeTemplate)(nil))
}
