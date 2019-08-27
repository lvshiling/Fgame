package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//仙桃大会常量模板配置
type XianTaoConstantTemplate struct {
	*XianTaoConstantTemplateVO
	thousandPeachBuffTemplate *BuffTemplate
	hundredPeachBuffTemplate  *BuffTemplate
	buSunPeachBuffTemplate    *BuffTemplate
	taoXianTemp               *BiologyTemplate
}

func (t *XianTaoConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.XianTaoBuff != 0 {
		buffTemplateVO := template.GetTemplateService().Get(int(t.XianTaoBuff), (*BuffTemplate)(nil))
		if buffTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", t.XianTaoBuff)
			err = template.NewTemplateFieldError("XianTaoBuff", err)
			return
		}
		t.thousandPeachBuffTemplate = buffTemplateVO.(*BuffTemplate)
	}

	if t.XianTaoBuff2 != 0 {
		buffTemplateVO := template.GetTemplateService().Get(int(t.XianTaoBuff2), (*BuffTemplate)(nil))
		if buffTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", t.XianTaoBuff2)
			err = template.NewTemplateFieldError("XianTaoBuff2", err)
			return
		}
		t.hundredPeachBuffTemplate = buffTemplateVO.(*BuffTemplate)
	}

	if t.BuSunBuff != 0 {
		buffTemplateVO := template.GetTemplateService().Get(int(t.BuSunBuff), (*BuffTemplate)(nil))
		if buffTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", t.BuSunBuff)
			err = template.NewTemplateFieldError("BuSunBuff", err)
			return
		}
		t.buSunPeachBuffTemplate = buffTemplateVO.(*BuffTemplate)
	}

	//桃仙id
	bilogyTemp := template.GetTemplateService().Get(int(t.TaoXianBiologyId), (*BiologyTemplate)(nil))
	if bilogyTemp == nil {
		err = fmt.Errorf("[%d] invalid", t.TaoXianBiologyId)
		return template.NewTemplateFieldError("TaoXianBiologyId", err)
	}
	taoXianTemp := bilogyTemp.(*BiologyTemplate)
	if taoXianTemp == nil {
		err = fmt.Errorf("[%d] invalid", t.TaoXianBiologyId)
		return template.NewTemplateFieldError("TaoXianBiologyId", err)
	}
	t.taoXianTemp = taoXianTemp

	return
}

func (t *XianTaoConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if err = validator.MinValidate(float64(t.XianTaoMin), float64(1), true); err != nil {
		err = template.NewTemplateFieldError("XianTaoMin", err)
		return
	}

	if err = validator.MinValidate(float64(t.XianTaoMax), float64(t.XianTaoMin), true); err != nil {
		err = template.NewTemplateFieldError("XianTaoMax", err)
		return
	}

	if err = validator.MinValidate(float64(t.CaiJiLimit), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("CaiJiLimit", err)
		return
	}

	if err = validator.MinValidate(float64(t.JieQuLimit), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("JieQuLimit", err)
		return
	}

	if err = validator.MinValidate(float64(t.TiJiaoService), float64(1), true); err != nil {
		err = template.NewTemplateFieldError("TiJiaoService", err)
		return
	}

	return
}

func (t *XianTaoConstantTemplate) PatchAfterCheck() {

}
func (t *XianTaoConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *XianTaoConstantTemplate) GetThousandPeachBuffTemplate() *BuffTemplate {
	return t.thousandPeachBuffTemplate
}

func (t *XianTaoConstantTemplate) GetHundredPeachBuffTemplate() *BuffTemplate {
	return t.hundredPeachBuffTemplate
}

func (t *XianTaoConstantTemplate) GetBuSunPeachBuffTemplate() *BuffTemplate {
	return t.buSunPeachBuffTemplate
}

func (t *XianTaoConstantTemplate) GetTaoXianTemp() *BiologyTemplate {
	return t.taoXianTemp
}

func (t *XianTaoConstantTemplate) FileName() string {
	return "tb_xiantao_constant.json"
}

func init() {
	template.Register((*XianTaoConstantTemplate)(nil))
}
