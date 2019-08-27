package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//开服循环活动时间配置
type OpenActivityXunHuanTimeTemplate struct {
	*OpenActivityXunHuanTimeTemplateVO              //
	timeTypeRange                      *randomGroup //类型范围
}

func (t *OpenActivityXunHuanTimeTemplate) TemplateId() int {
	return t.Id
}

func (t *OpenActivityXunHuanTimeTemplate) IsOnRange(openDay int32) bool {
	if openDay >= t.timeTypeRange.min && openDay <= t.timeTypeRange.max {
		return true
	}
	return false
}

func (t *OpenActivityXunHuanTimeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.timeTypeRange = &randomGroup{
		min: t.OpentimeBegin,
		max: t.OpentimeFinish,
	}
	return nil
}

func (t *OpenActivityXunHuanTimeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	err = validator.MinValidate(float64(t.OpentimeBegin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.OpentimeBegin)
		err = template.NewTemplateFieldError("OpentimeBegin", err)
		return
	}

	err = validator.MinValidate(float64(t.OpentimeFinish), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.OpentimeFinish)
		err = template.NewTemplateFieldError("OpentimeFinish", err)
		return
	}

	//
	if t.timeTypeRange.min > t.timeTypeRange.max {
		err = fmt.Errorf("[%d][%d] invalid", t.OpentimeBegin, t.OpentimeFinish)
		err = template.NewTemplateFieldError("OpentimeBegin or OpentimeFinish", err)
		return
	}

	// if t.nextTemp != nil {
	// 	//区间连续校验
	// 	if t.nextTemp.tonicRange.min-t.tonicRange.max != 1 {
	// 		err = fmt.Errorf("[%d] invalid", t.QuJian)
	// 		return template.NewTemplateFieldError("QuJian", err)
	// 	}
	// }

	return nil
}

func (t *OpenActivityXunHuanTimeTemplate) PatchAfterCheck() {

}

func (t *OpenActivityXunHuanTimeTemplate) FileName() string {
	return "tb_circle_type.json"
}

func init() {
	template.Register((*OpenActivityXunHuanTimeTemplate)(nil))
}
