package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//活动怪物掉落配置
type TeShuDropTemplate struct {
	*TeShuDropTemplateVO
}

func (t *TeShuDropTemplate) TemplateId() int {
	return t.Id
}

func (t *TeShuDropTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *TeShuDropTemplate) PatchAfterCheck() {

}

func (t *TeShuDropTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//掉落id
	if err = validator.MinValidate(float64(t.DropId), float64(0), false); err != nil {
		err = template.NewTemplateFieldError("DropId", err)
		return
	}

	// 区间
	if t.MinCount > t.MaxCount {
		err = fmt.Errorf("[%d] invalid", t.MinCount)
		err = template.NewTemplateFieldError("MinCount", err)
	}

	// 活动id
	if err = validator.MinValidate(float64(t.GroupId), float64(0), false); err != nil {
		err = template.NewTemplateFieldError("GroupId", err)
		return
	}

	return nil
}

func (t *TeShuDropTemplate) FileName() string {
	return "tb_teshu_drop.json"
}

func init() {
	template.Register((*TeShuDropTemplate)(nil))
}
