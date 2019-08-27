package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

// 特戒进阶配置
type RingAdvanceTemplate struct {
	*RingAdvanceTemplateVO
	nextTemp    *RingAdvanceTemplate
	needItemMap map[int32]int32
}

func (t *RingAdvanceTemplate) TemplateId() int {
	return t.Id
}

func (t *RingAdvanceTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *RingAdvanceTemplate) IsClearBless() bool {
	return t.IsClear == 1
}

func (t *RingAdvanceTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.needItemMap = make(map[int32]int32)
	//验证物品id
	if t.UseItem != 0 {
		to := template.GetTemplateService().Get(int(t.UseItem), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}

		err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}
		t.needItemMap[t.UseItem] = t.ItemCount
	}

	return
}

func (t *RingAdvanceTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*RingAdvanceTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*RingAdvanceTemplate)
		diff := t.nextTemp.Advance - int32(t.Advance)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	// 验证升阶成功率
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		return template.NewTemplateFieldError("UpdateWfb", err)
	}

	//验证 hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//验证 defence
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(t.TimesMin), float64(0), true, float64(t.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(t.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(t.AddMin), float64(0), true, float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(t.AddMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 is_clear
	if t.IsClear != 0 && t.IsClear != 1 {
		err = fmt.Errorf("[%d] invalid", t.IsClear)
		return template.NewTemplateFieldError("IsClear", err)
	}

	return nil
}

func (t *RingAdvanceTemplate) PatchAfterCheck() {
}

func (t *RingAdvanceTemplate) FileName() string {
	return "tb_tejie_jinjie.json"
}

func init() {
	template.Register((*RingAdvanceTemplate)(nil))
}
