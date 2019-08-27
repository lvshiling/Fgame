package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//元神金装开光配置
type GoldEquipOpenLightTemplate struct {
	*GoldEquipOpenLightTemplateVO
	nextTemp        *GoldEquipOpenLightTemplate
	needItemMap     map[int32]int32
	tunshiReturnMap map[int32]int32
}

func (t *GoldEquipOpenLightTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldEquipOpenLightTemplate) GetNextTemplate() *GoldEquipOpenLightTemplate {
	return t.nextTemp
}

func (t *GoldEquipOpenLightTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *GoldEquipOpenLightTemplate) GetReturnItemMap() map[int32]int32 {
	return t.tunshiReturnMap
}

func (t *GoldEquipOpenLightTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//开光物品
	t.needItemMap = make(map[int32]int32)
	needItemIdList, err := utils.SplitAsIntArray(t.UseItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItem)
		return template.NewTemplateFieldError("UseItem", err)
	}
	needItemCountList, err := utils.SplitAsIntArray(t.UseCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseCount)
		return template.NewTemplateFieldError("UseCount", err)
	}
	if len(needItemIdList) != len(needItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.UseItem, t.UseCount)
		err = template.NewTemplateFieldError("UseItem or UseCount", err)
		return err
	}
	if len(needItemIdList) > 0 {
		//组合数据
		for index, itemId := range needItemIdList {
			_, ok := t.needItemMap[itemId]
			if ok {
				t.needItemMap[itemId] += needItemCountList[index]
			} else {
				t.needItemMap[itemId] = needItemCountList[index]
			}
		}
	}

	// 吞噬返还
	t.tunshiReturnMap = make(map[int32]int32)
	returnItemIdList, err := utils.SplitAsIntArray(t.MeltingReturnId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MeltingReturnId)
		return template.NewTemplateFieldError("MeltingReturnId", err)
	}
	returnItemCountList, err := utils.SplitAsIntArray(t.MeltingReturnCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MeltingReturnCount)
		return template.NewTemplateFieldError("MeltingReturnCount", err)
	}
	if len(returnItemIdList) != len(returnItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.MeltingReturnId, t.MeltingReturnCount)
		err = template.NewTemplateFieldError("MeltingReturnId or MeltingReturnCount", err)
		return err
	}
	if len(returnItemIdList) > 0 {
		//组合数据
		for index, itemId := range returnItemIdList {
			_, ok := t.tunshiReturnMap[itemId]
			if ok {
				t.tunshiReturnMap[itemId] += returnItemCountList[index]
			} else {
				t.tunshiReturnMap[itemId] = returnItemCountList[index]
			}
		}
	}

	return nil
}

func (t *GoldEquipOpenLightTemplate) PatchAfterCheck() {
}

func (t *GoldEquipOpenLightTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*GoldEquipOpenLightTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextTemp = to.(*GoldEquipOpenLightTemplate)
	}

	//开光次数
	if err = validator.MinValidate(float64(t.Times), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("Times", err)
		return
	}

	//成功率
	if err = validator.MinValidate(float64(t.SuccessRate), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("SuccessRate", err)
		return
	}
	//最小次数
	if err = validator.MinValidate(float64(t.TimesMin), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}
	//最大次数
	if err = validator.MinValidate(float64(t.TimesMax), float64(0), true); err != nil {
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

	//属性加成
	if err = validator.MinValidate(float64(t.AttrPercent), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("AttrPercent", err)
		return
	}

	// 消耗物品id
	for itemId, num := range t.needItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("UseItem", fmt.Errorf("[%d] invalid", itemId))
		}
		if err = validator.MinValidate(float64(num), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("UseCount", err)
			return
		}
	}

	// 返回数量
	for itemId, num := range t.tunshiReturnMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("MeltingReturnId", fmt.Errorf("[%d] invalid", itemId))
		}
		if err = validator.MinValidate(float64(num), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("MeltingReturnCount", err)
			return
		}
	}

	return nil
}

func (t *GoldEquipOpenLightTemplate) FileName() string {
	return "tb_goldequip_openlight.json"
}

func init() {
	template.Register((*GoldEquipOpenLightTemplate)(nil))
}
