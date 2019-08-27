package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type GodCastingJichengTemplate struct {
	*GodCastingJichengTemplateVO
}

func (t *GodCastingJichengTemplate) TemplateId() int {
	return t.Id
}

func (t *GodCastingJichengTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}
func (t *GodCastingJichengTemplate) PatchAfterCheck() {

}
func (t *GodCastingJichengTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Level", err)
	}

	//物品id
	itemTmpObj := template.GetTemplateService().Get(int(t.NeedItemCount), (*ItemTemplate)(nil))
	if itemTmpObj == nil {
		return template.NewTemplateFieldError("NeedItemId", fmt.Errorf("[%d] invalid", t.NeedItemCount))
	}

	//物品数量
	err = validator.MinValidate(float64(t.NeedItemCount), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("NeedItemCount", err)
	}

	return nil
}

func (edt *GodCastingJichengTemplate) FileName() string {
	return "tb_shenzhu_jicheng.json"
}

func init() {
	template.Register((*GodCastingJichengTemplate)(nil))
}
