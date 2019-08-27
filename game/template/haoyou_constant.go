package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//推送常量配置
type NoticeConstantTemplate struct {
	*NoticeConstantTemplateVO
}

func (t *NoticeConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *NoticeConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *NoticeConstantTemplate) PatchAfterCheck() {
}

func (t *NoticeConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//购买所需元宝
	err = validator.MinValidate(float64(t.BaoHuFei), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BaoHuFei)
		return template.NewTemplateFieldError("BaoHuFei", err)
	}

	// 持续时间
	err = validator.MinValidate(float64(t.BaoHuTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BaoHuTime)
		return template.NewTemplateFieldError("BaoHuTime", err)
	}

	// 损失银两
	err = validator.MinValidate(float64(t.ChouRenSunShiSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ChouRenSunShiSilver)
		return template.NewTemplateFieldError("ChouRenSunShiSilver", err)
	}

	// 赞赏奖励次数
	err = validator.MinValidate(float64(t.ShouliLimitCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ShouliLimitCount)
		return template.NewTemplateFieldError("ShouliLimitCount", err)
	}

	return nil
}

func (edt *NoticeConstantTemplate) FileName() string {
	return "tb_haoyou_constant.json"
}

func init() {
	template.Register((*NoticeConstantTemplate)(nil))
}
