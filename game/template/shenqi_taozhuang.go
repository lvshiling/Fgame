package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//神器套装配置
type ShenQiTaoZhuangTemplate struct {
	*ShenQiTaoZhuangTemplateVO
}

func (est *ShenQiTaoZhuangTemplate) TemplateId() int {
	return est.Id
}

func (et *ShenQiTaoZhuangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	return nil
}
func (et *ShenQiTaoZhuangTemplate) PatchAfterCheck() {

}
func (et *ShenQiTaoZhuangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//穿戴的器灵最低阶数
	err = validator.MinValidate(float64(et.NeedNumber), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.NeedNumber)
		return template.NewTemplateFieldError("NeedNumber", err)
	}

	//达到要求的器灵数量
	err = validator.MinValidate(float64(et.NeedCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.NeedCount)
		return template.NewTemplateFieldError("NeedCount", err)
	}

	//生命
	err = validator.MinValidate(float64(et.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}
	//攻击
	err = validator.MinValidate(float64(et.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}
	//防御
	err = validator.MinValidate(float64(et.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	return nil
}

func (edt *ShenQiTaoZhuangTemplate) FileName() string {
	return "tb_shenqi_taozhuang.json"
}

func init() {
	template.Register((*ShenQiTaoZhuangTemplate)(nil))
}
