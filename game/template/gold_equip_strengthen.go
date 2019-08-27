package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//装备强化配置
type GoldEquipStrengthenTemplate struct {
	*GoldEquipStrengthenTemplateVO
	rateMap                         map[int32]int32              //提供强化概率
	nextGoldEquipStrengthenTemplate *GoldEquipStrengthenTemplate //下一级强化
	tunshiReturnMap                 map[int32]int32
	needItemList                    []int32
}

func (est *GoldEquipStrengthenTemplate) TemplateId() int {
	return est.Id
}

func (t *GoldEquipStrengthenTemplate) GetReturnItemMap() map[int32]int32 {
	return t.tunshiReturnMap
}

func (t *GoldEquipStrengthenTemplate) GetNeedItemMap(needNum int32) map[int32]int32 {
	itemMap := make(map[int32]int32)
	for _, itemId := range t.needItemList {
		itemMap[itemId] = needNum
	}

	return itemMap
}

func (t *GoldEquipStrengthenTemplate) OfferRate(level int32) int32 {
	return t.rateMap[level]
}

func (t *GoldEquipStrengthenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//下一阶强化
	if t.NextId != 0 {
		tempNextGoldEquipStrengthenTemplate := template.GetTemplateService().Get(int(t.NextId), (*GoldEquipStrengthenTemplate)(nil))
		if tempNextGoldEquipStrengthenTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}
		t.nextGoldEquipStrengthenTemplate = tempNextGoldEquipStrengthenTemplate.(*GoldEquipStrengthenTemplate)
	}

	//提供的概率
	t.rateMap = make(map[int32]int32)
	if t.GiveRate1 != 0 {
		t.rateMap[0] = t.GiveRate1
	}
	if t.GiveRate2 != 0 {
		t.rateMap[1] = t.GiveRate2
	}
	if t.GiveRate3 != 0 {
		t.rateMap[2] = t.GiveRate3
	}
	if t.GiveRate4 != 0 {
		t.rateMap[3] = t.GiveRate4
	}
	if t.GiveRate5 != 0 {
		t.rateMap[4] = t.GiveRate5
	}
	if t.GiveRate6 != 0 {
		t.rateMap[5] = t.GiveRate6
	}
	if t.GiveRate7 != 0 {
		t.rateMap[6] = t.GiveRate7
	}
	if t.GiveRate8 != 0 {
		t.rateMap[7] = t.GiveRate8
	}
	if t.GiveRate9 != 0 {
		t.rateMap[8] = t.GiveRate9
	}
	if t.GiveRate10 != 0 {
		t.rateMap[9] = t.GiveRate10
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
		return template.NewTemplateFieldError("MeltingReturnId or MeltingReturnCount", err)
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

	// 强化需要的物品
	needItemIdList, err := utils.SplitAsIntArray(t.NeedItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MeltingReturnId)
		return template.NewTemplateFieldError("MeltingReturnId", err)
	}
	t.needItemList = needItemIdList

	return nil
}
func (t *GoldEquipStrengthenTemplate) PatchAfterCheck() {

}
func (t *GoldEquipStrengthenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	for itemId, num := range t.tunshiReturnMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			err = fmt.Errorf("[%s] invalid", t.MeltingReturnId)
			return template.NewTemplateFieldError("MeltingReturnId", err)
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.MeltingReturnCount)
			return template.NewTemplateFieldError("MeltingReturnCount", err)
		}

	}

	return nil
}

func (edt *GoldEquipStrengthenTemplate) FileName() string {
	return "tb_goldequip_strengthen.json"
}

func init() {
	template.Register((*GoldEquipStrengthenTemplate)(nil))
}
