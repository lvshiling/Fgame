package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//寻宝次数配置
type HuntTimesTemplate struct {
	*HuntTimesTemplateVO
	nextTemp *HuntTimesTemplate
}

func (t *HuntTimesTemplate) TemplateId() int {
	return t.Id
}

func (t *HuntTimesTemplate) GetNextTemp() *HuntTimesTemplate {
	return t.nextTemp
}

func (t *HuntTimesTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return
}

func (t *HuntTimesTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// nextId
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("JianGeTime", err)
		}

		to := template.GetTemplateService().Get(int(t.NextId), (*HuntTimesTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("JianGeTime", err)
		}

		t.nextTemp = to.(*HuntTimesTemplate)
	}

	// 次数
	err = validator.MinValidate(float64(t.Times), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Times", err)
	}
	// 时间间隔
	err = validator.MinValidate(float64(t.JianGeTime), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("JianGeTime", err)
	}

	return nil
}

func (t *HuntTimesTemplate) PatchAfterCheck() {
}

func (t *HuntTimesTemplate) FileName() string {
	return "tb_xunbao_times.json"
}

func init() {
	template.Register((*HuntTimesTemplate)(nil))
}
