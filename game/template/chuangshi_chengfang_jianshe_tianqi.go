package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type ChuangShiChengFangJianSheTianQiTemplate struct {
	*ChuangShiChengFangJianSheTianQiTemplateVO
	needItemMap map[int32]int32
}

func (t *ChuangShiChengFangJianSheTianQiTemplate) GetActivateItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *ChuangShiChengFangJianSheTianQiTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiChengFangJianSheTianQiTemplate) FileName() string {
	return "tb_chuangshi_tianqi.json"
}

func (t *ChuangShiChengFangJianSheTianQiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return
}

func (t *ChuangShiChengFangJianSheTianQiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	err = validator.MinValidate(float64(t.ValidTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ValidTime)
		return template.NewTemplateFieldError("ValidTime", err)
	}

	to := template.GetTemplateService().Get(int(t.TianqiItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.TianqiItemId)
		err = template.NewTemplateFieldError("TianqiItemId", err)
		return
	}
	//数量
	err = validator.MinValidate(float64(t.TianqiItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TianqiItemCount)
		return template.NewTemplateFieldError("TianqiItemCount", err)
	}

	return
}

func (t *ChuangShiChengFangJianSheTianQiTemplate) PatchAfterCheck() {
	t.needItemMap = make(map[int32]int32)
	t.needItemMap[t.TianqiItemId] = t.TianqiItemCount
}

func init() {
	template.Register((*ChuangShiChengFangJianSheTianQiTemplate)(nil))
}
