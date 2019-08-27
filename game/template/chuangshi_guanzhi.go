package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

type ChuangShiGuanZhiTemplate struct {
	*ChuangShiGuanZhiTemplateVO
	useItemMap     map[int32]int32
	receiveItemMap map[int32]int32
	fashionMap     map[int32]int32
	nextTemplate   *ChuangShiGuanZhiTemplate
}

func (t *ChuangShiGuanZhiTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiGuanZhiTemplate) GetReceiveItemMap() map[int32]int32 {
	return t.receiveItemMap
}

func (t *ChuangShiGuanZhiTemplate) GetUseItemMap() map[int32]int32 {
	return t.useItemMap
}

func (t *ChuangShiGuanZhiTemplate) GetFashionMap() map[int32]int32 {
	return t.fashionMap
}

func (t *ChuangShiGuanZhiTemplate) FileName() string {
	return "tb_chuangshi_guanzhi.json"
}

func (t *ChuangShiGuanZhiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.fashionMap = make(map[int32]int32)

	// 升级消耗
	t.useItemMap = make(map[int32]int32)
	if t.ItemId != 0 {
		to := template.GetTemplateService().Get(int(t.ItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.ItemId)
			return template.NewTemplateFieldError("ItemId", err)
		}

		err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}
		t.useItemMap[t.ItemId] = t.ItemCount
	}

	// 升级获得
	t.receiveItemMap = make(map[int32]int32)
	if t.GetItemId != 0 {
		to := template.GetTemplateService().Get(int(t.GetItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.GetItemId)
			return template.NewTemplateFieldError("GetItemId", err)
		}

		err = validator.MinValidate(float64(t.GetItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.GetItemCount)
			return template.NewTemplateFieldError("GetItemCount", err)
		}
		t.receiveItemMap[t.GetItemId] = t.GetItemCount
	}

	return
}

func (t *ChuangShiGuanZhiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.NextId != 0 {
		temp := template.GetTemplateService().Get(t.NextId, (*ChuangShiGuanZhiTemplate)(nil))
		guanZhiTemp, ok := temp.(*ChuangShiGuanZhiTemplate)
		if !ok {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemplate = guanZhiTemp
		if t.nextTemplate.Level-t.Level != 1 {
			err = fmt.Errorf("[%d] invalid", t.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	// 验证升级概率
	err = validator.RangeValidate(float64(t.UpLevPercent), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpLevPercent)
		return template.NewTemplateFieldError("UpLevPercent", err)
	}
	// 验证使用威望
	err = validator.MinValidate(float64(t.UseWeiWang), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseWeiWang)
		return template.NewTemplateFieldError("UseWeiWang", err)
	}

	// 验证使用银两
	err = validator.MinValidate(float64(t.UseMoney), 0, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseMoney)
		return template.NewTemplateFieldError("UseMoney", err)
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

	// 验证时装id
	if t.FashionId != 0 {
		temp := template.GetTemplateService().Get(int(t.FashionId), (*FashionTemplate)(nil))
		if temp == nil {
			err = fmt.Errorf("[%d] invalid", t.FashionId)
			err = template.NewTemplateFieldError("FashionId", err)
			return
		}
		t.fashionMap[t.FashionId] = 1
	}
	return
}

func (t *ChuangShiGuanZhiTemplate) PatchAfterCheck() {
}

func init() {
	template.Register((*ChuangShiGuanZhiTemplate)(nil))
}
