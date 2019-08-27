package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

const (
	maxTotalTrade = 1000
)

//交易市场配置
type TradeConstantTemplate struct {
	*TradeConstantTemplateVO
}

func (t *TradeConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *TradeConstantTemplate) PatchAfterCheck() {
}

func (t *TradeConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *TradeConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//总的上架数量
	err = validator.MaxValidate(float64(t.AllCountMax), float64(maxTotalTrade), true)
	if err != nil {
		return template.NewTemplateFieldError("AllCountMax", fmt.Errorf("[%d] invalid", t.AllCountMax))
	}

	return nil
}

func (t *TradeConstantTemplate) FileName() string {
	return "tb_jiaoyishichang_constant.json"
}

func init() {
	template.Register((*TradeConstantTemplate)(nil))
}
