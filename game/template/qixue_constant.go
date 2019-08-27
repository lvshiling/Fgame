package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

//泣血枪常量配置
type QiXueConstantTemplate struct {
	*QiXueConstantTemplateVO
}

func (t *QiXueConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *QiXueConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *QiXueConstantTemplate) PatchAfterCheck() {

}

func (t *QiXueConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//DropCd
	err = validator.MinValidate(float64(t.DropRate), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropRate)
		return template.NewTemplateFieldError("DropRate", err)
	}

	//KeyDropPercent
	err = validator.MinValidate(float64(t.DropPercentMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropPercentMin)
		return template.NewTemplateFieldError("DropPercentMin", err)
	}
	//KeyDropPercent
	err = validator.RangeValidate(float64(t.DropPercentMax), float64(t.DropPercentMin), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropPercentMax)
		return template.NewTemplateFieldError("DropPercentMax", err)
	}

	//DropCd
	err = validator.MinValidate(float64(t.DropCd), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropCd)
		return template.NewTemplateFieldError("DropCd", err)
	}

	//protected_time
	err = validator.MinValidate(float64(t.DropProtectedTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropProtectedTime)
		return template.NewTemplateFieldError("DropProtectedTime", err)
	}

	//fail_time
	err = validator.MinValidate(float64(t.DropFailTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropFailTime)
		return template.NewTemplateFieldError("DropFailTime", err)
	}

	//min_stack
	err = validator.MinValidate(float64(t.DropMinStack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropMinStack)
		return template.NewTemplateFieldError("DropMinStack", err)
	}

	//max_stack
	err = validator.MinValidate(float64(t.DropMaxStack), float64(t.DropMinStack), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropMaxStack)
		return template.NewTemplateFieldError("DropMaxStack", err)
	}

	//min_stack
	err = validator.RangeValidate(float64(t.DropSystemReturn), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropSystemReturn)
		return template.NewTemplateFieldError("DropSystemReturn", err)
	}

	return nil
}

func (t *QiXueConstantTemplate) FileName() string {
	return "tb_qixue_constant.json"
}

func init() {
	template.Register((*QiXueConstantTemplate)(nil))
}
