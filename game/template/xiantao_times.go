package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

//仙桃大会劫取次数配置
type XianTaoTimesTemplate struct {
	*XianTaoTimesTemplateVO
	nextTemp *XianTaoTimesTemplate //下一条
}

func (mclt *XianTaoTimesTemplate) TemplateId() int {
	return mclt.Id
}

func (mclt *XianTaoTimesTemplate) GetNextTemplate() *XianTaoTimesTemplate {
	return mclt.nextTemp
}

func (mclt *XianTaoTimesTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//下一条
	if mclt.NextId != 0 {
		if mclt.NextId-mclt.Id != 1 {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		to := template.GetTemplateService().Get(mclt.NextId, (*XianTaoTimesTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		mclt.nextTemp = to.(*XianTaoTimesTemplate)

	}

	return nil
}

func (mclt *XianTaoTimesTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//验证 times
	err = validator.MinValidate(float64(mclt.Times), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Times)
		err = template.NewTemplateFieldError("Times", err)
		return
	}

	if mclt.nextTemp != nil {
		//区间连续校验
		if mclt.nextTemp.Times-mclt.Times != 1 {
			err = fmt.Errorf("[%d] invalid", mclt.Times)
			return template.NewTemplateFieldError("Times", err)
		}
	}

	//验证 JieQuPercent
	err = validator.RangeValidate(float64(mclt.JieQuPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.JieQuPercent)
		err = template.NewTemplateFieldError("JieQuPercent", err)
		return
	}

	//验证 SunShiPercent
	err = validator.RangeValidate(float64(mclt.SunShiPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.SunShiPercent)
		err = template.NewTemplateFieldError("SunShiPercent", err)
		return
	}

	return nil
}

func (mclt *XianTaoTimesTemplate) PatchAfterCheck() {

}

func (mclt *XianTaoTimesTemplate) FileName() string {
	return "tb_xiantao_times.json"
}

func init() {
	template.Register((*XianTaoTimesTemplate)(nil))
}
