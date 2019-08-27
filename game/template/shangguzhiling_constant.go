package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type ShangguzhilingConstantTemplate struct {
	*ShangguzhilingConstantTemplateVO
}

func (t *ShangguzhilingConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingConstantTemplate) PatchAfterCheck() {

}

func (t *ShangguzhilingConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//上古之灵灵炼锁定使用的物品ID
	err = validator.MinValidate(float64(t.LinglianSuodingUseItemId), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LinglianSuodingUseItemId)
		return template.NewTemplateFieldError("LinglianSuodingUseItemId", err)
	}

	//灵炼消耗物品ID
	err = validator.MinValidate(float64(t.LinglianItemId), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LinglianItemId)
		return template.NewTemplateFieldError("LinglianItemId", err)
	}
	//灵炼消耗物品数量基数
	err = validator.MinValidate(float64(t.LinglianItemCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LinglianItemCount)
		return template.NewTemplateFieldError("LinglianItemCount", err)
	}
	//灵炼消耗物品数量参数1
	err = validator.MinValidate(float64(t.LinglianCoefficient), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LinglianCoefficient)
		return template.NewTemplateFieldError("LinglianCoefficient", err)
	}
	//灵炼消耗物品数量参数2
	err = validator.MinValidate(float64(t.LinglianCoefficient2), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LinglianCoefficient2)
		return template.NewTemplateFieldError("LinglianCoefficient2", err)
	}
	//灵炼次数上限
	err = validator.MinValidate(float64(t.LinglianNum), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LinglianNum)
		return template.NewTemplateFieldError("LinglianNum", err)
	}

	return nil
}

func (t *ShangguzhilingConstantTemplate) FileName() string {
	return "tb_sgzl_constant.json"
}

func init() {
	template.Register((*ShangguzhilingConstantTemplate)(nil))
}
