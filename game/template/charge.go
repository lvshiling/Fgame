package template

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//充值配置
type ChargeTemplate struct {
	*ChargeTemplateVO
	typ logintypes.SDKType
}

func (t *ChargeTemplate) TemplateId() int {
	return t.Id
}

func (t *ChargeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.typ = logintypes.SDKType(t.Type)

	return
}

func (t *ChargeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	if !t.typ.Valid() {
		err = fmt.Errorf("%d invalid", t.Type)
		return template.NewTemplateFieldError("type", err)
	}

	// 服务器排序
	err = validator.MinValidate(float64(t.SubType), float64(1), true)
	if err != nil {
		err = fmt.Errorf("%d invalid", t.SubType)
		return template.NewTemplateFieldError("SubType", err)
	}
	// 金额
	err = validator.MinValidate(float64(t.Rmb), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Rmb", err)
	}
	// 元宝
	err = validator.MinValidate(float64(t.Gold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Gold", err)
	}
	// 首次返还
	err = validator.MinValidate(float64(t.FanhuanBindGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("FanhuanBindGold", err)
	}
	// 首次返还
	err = validator.MinValidate(float64(t.FanhuanGold), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("FanhuanGold", err)
	}

	return nil
}

func (t *ChargeTemplate) PatchAfterCheck() {
}

func (t *ChargeTemplate) GetType() logintypes.SDKType {
	return t.typ
}

func (t *ChargeTemplate) FileName() string {
	return "tb_recharge.json"
}

func init() {
	template.Register((*ChargeTemplate)(nil))
}
