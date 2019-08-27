package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//折扣配置
type DiscountTemplate struct {
	*DiscountTemplateVO
}

func (t *DiscountTemplate) TemplateId() int {
	return t.Id
}

func (t *DiscountTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *DiscountTemplate) PatchAfterCheck() {

}

func (t *DiscountTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 buyCount
	err = validator.MinValidate(float64(t.BuyCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BuyCount)
		err = template.NewTemplateFieldError("BuyCount", err)
		return
	}

	//验证 maxCount
	if t.MaxCount < t.BuyCount {
		err = fmt.Errorf("[%d] invalid", t.BuyCount)
		err = template.NewTemplateFieldError("MaxCount", err)
		return
	}

	// 原价
	err = validator.MinValidate(float64(t.YuanGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("YuanGold", err)
	}

	//现价
	err = validator.MinValidate(float64(t.UseGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("UseGold", err)
	}

	return nil
}

func (t *DiscountTemplate) FileName() string {
	return "tb_xianshi_gift.json"
}

func init() {
	template.Register((*DiscountTemplate)(nil))
}
