package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

//神域常量配置
type ShenYuConstantTemplate struct {
	*ShenYuConstantTemplateVO
}

func (t *ShenYuConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *ShenYuConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShenYuConstantTemplate) PatchAfterCheck() {

}

func (t *ShenYuConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//key_max
	err = validator.MinValidate(float64(t.KeyMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KeyMax)
		return template.NewTemplateFieldError("KeyMax", err)
	}
	//key_max
	err = validator.MaxValidate(float64(t.KeyMax), float64(255), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KeyMax)
		return template.NewTemplateFieldError("KeyMax", err)
	}

	//KeyKeepMin
	err = validator.MinValidate(float64(t.KeyKeepMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KeyKeepMin)
		return template.NewTemplateFieldError("KeyKeepMin", err)
	}

	//KeyDropPercent
	err = validator.MinValidate(float64(t.KeyDropPercent), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KeyDropPercent)
		return template.NewTemplateFieldError("KeyDropPercent", err)
	}
	//KeyDropPercent
	err = validator.MaxValidate(float64(t.KeyDropPercent), float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KeyDropPercent)
		return template.NewTemplateFieldError("KeyDropPercent", err)
	}

	//KeyDropCD
	err = validator.MinValidate(float64(t.KeyDropCD), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KeyDropCD)
		return template.NewTemplateFieldError("KeyDropCD", err)
	}

	//min_stack
	err = validator.MinValidate(float64(t.MinStack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MinStack)
		return template.NewTemplateFieldError("MinStack", err)
	}

	//max_stack
	err = validator.MinValidate(float64(t.MaxStack), float64(t.MinStack), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MaxStack)
		return template.NewTemplateFieldError("MaxStack", err)
	}

	//exist_time
	err = validator.MinValidate(float64(t.ExistTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ExistTime)
		return template.NewTemplateFieldError("ExistTime", err)
	}

	//protected_time
	err = validator.MinValidate(float64(t.ProtectedTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ProtectedTime)
		return template.NewTemplateFieldError("ProtectedTime", err)
	}

	//fail_time
	err = validator.MinValidate(float64(t.FailTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FailTime)
		return template.NewTemplateFieldError("FailTime", err)
	}

	return nil
}

func (t *ShenYuConstantTemplate) FileName() string {
	return "tb_shenyu_constant.json"
}

func init() {
	template.Register((*ShenYuConstantTemplate)(nil))
}
