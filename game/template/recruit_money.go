package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//运营活动-元宝金猪配置
type GoldPigTemplate struct {
	*GoldPigTemplateVO
	nextTemp *GoldPigTemplate
}

func (t *GoldPigTemplate) GetNextTemplate() *GoldPigTemplate {
	return t.nextTemp
}

func (t *GoldPigTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldPigTemplate) PatchAfterCheck() {
}

func (t *GoldPigTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *GoldPigTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//nextId
	//验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - t.Id
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(t.NextId, (*GoldPigTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*GoldPigTemplate)
	}

	//验证 活动id
	err = validator.MinValidate(float64(t.Group), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Group)
		err = template.NewTemplateFieldError("Group", err)
		return
	}

	//验证 充值
	err = validator.MinValidate(float64(t.NeedRecharge), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedRecharge)
		err = template.NewTemplateFieldError("NeedRecharge", err)
		return
	}

	//验证 返回比例
	err = validator.MinValidate(float64(t.ReturnPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ReturnPercent)
		err = template.NewTemplateFieldError("ReturnPercent", err)
		return
	}

	return nil
}

func (t *GoldPigTemplate) FileName() string {
	return "tb_recruit_money.json"
}

func init() {
	template.Register((*GoldPigTemplate)(nil))
}
