package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//圣坛配置
type UnionShengTanTemplate struct {
	*UnionShengTanTemplateVO
}

func (t *UnionShengTanTemplate) TemplateId() int {
	return t.Id
}

func (t *UnionShengTanTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//FirstTime
	err = validator.MinValidate(float64(t.FirstTime), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("FirstTime", err)
		return
	}

	//RewTime
	err = validator.MinValidate(float64(t.RewTime), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewTime", err)
		return
	}

	//ExpAddItemLimit
	err = validator.MinValidate(float64(t.ExpAddItemLimit), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("ExpAddItemLimit", err)
		return
	}

	//XiaoguaiTime
	err = validator.MinValidate(float64(t.XiaoguaiTime), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("XiaoguaiTime", err)
		return
	}

	return nil
}

const (
	minThreshold = 100
)

func (t *UnionShengTanTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	tempBiologyTemplate := template.GetTemplateService().Get(int(t.ShengtanId), (*BiologyTemplate)(nil))
	if tempBiologyTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.ShengtanId)
		err = template.NewTemplateFieldError("ShengtanId", err)
		return
	}

	//仙盟最高人数
	err = validator.MinValidate(float64(t.GonggaoHpPercent), float64(minThreshold), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GonggaoHpPercent)
		return template.NewTemplateFieldError("GonggaoHpPercent", err)
	}

	return nil
}

func (t *UnionShengTanTemplate) PatchAfterCheck() {

}

func (t *UnionShengTanTemplate) FileName() string {
	return "tb_union_shengtan.json"
}

func init() {
	template.Register((*UnionShengTanTemplate)(nil))
}
