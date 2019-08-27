package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type ShangguzhilingJinjieTemplate struct {
	*ShangguzhilingJinjieTemplateVO
	nextLevelTemp *ShangguzhilingJinjieTemplate
}

func (t *ShangguzhilingJinjieTemplate) GetNextRankTemp() *ShangguzhilingJinjieTemplate {
	return t.nextLevelTemp
}

func (t *ShangguzhilingJinjieTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingJinjieTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingJinjieTemplate) PatchAfterCheck() {

}

func (t *ShangguzhilingJinjieTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//等级
	err = validator.MinValidate(float64(t.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Number)
		return template.NewTemplateFieldError("Number", err)
	}

	//下一等级模板
	if t.NextId != 0 {
		nextTempInterface := template.GetTemplateService().Get(int(t.NextId), (*ShangguzhilingJinjieTemplate)(nil))
		if nextTempInterface == nil {
			err = fmt.Errorf("ShangguzhilingJinjieTemplate [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp, ok := nextTempInterface.(*ShangguzhilingJinjieTemplate)
		if !ok {
			err = fmt.Errorf("ShangguzhilingJinjieTemplate assert [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		if nextTemp.Number != t.Number+1 {
			err = fmt.Errorf("ShangguzhilingJinjieTemplate [%d] invalid, curLevel [%d], nextTempLevel [%d]", t.NextId, t.Number, nextTemp.Number)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextLevelTemp = nextTemp
	}

	//消耗物品Id
	err = validator.MinValidate(float64(t.UseItem), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItem)
		return template.NewTemplateFieldError("UseItem", err)
	}

	//消耗物品数量
	err = validator.MinValidate(float64(t.ItemCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ItemCount)
		return template.NewTemplateFieldError("ItemCount", err)
	}

	//hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//攻击
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//防御
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//升级成功率万分比
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(10000), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		return template.NewTemplateFieldError("UpdateWfb", err)
	}

	//最小次数
	err = validator.MinValidate(float64(t.TimesMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		return template.NewTemplateFieldError("TimesMin", err)
	}

	//最大次数
	err = validator.MinValidate(float64(t.TimesMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		return template.NewTemplateFieldError("TimesMax", err)
	}

	//祝福值随机最小值
	err = validator.MinValidate(float64(t.AddMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMin)
		return template.NewTemplateFieldError("AddMin", err)
	}

	//祝福值随机最大值
	err = validator.MinValidate(float64(t.AddMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMax)
		return template.NewTemplateFieldError("AddMax", err)
	}

	//最大祝福值
	err = validator.MinValidate(float64(t.ZhufuMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		return template.NewTemplateFieldError("ZhufuMax", err)
	}

	return nil
}

func (t *ShangguzhilingJinjieTemplate) FileName() string {
	return "tb_sgzl_jinjie.json"
}

func init() {
	template.Register((*ShangguzhilingJinjieTemplate)(nil))
}
