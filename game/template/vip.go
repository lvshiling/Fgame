package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//VIP配置
type VipTemplate struct {
	*VipTemplateVO
	nextTemp             *VipTemplate
	disCountItemMap      map[int32]int32                            //折扣物品
	freeGiftItemMap      map[int32]int32                            //免费礼包物品
	emailDiscountItemMap map[int32]int32                            //邮件/展示
	emailFreeGiftItemMap map[int32]int32                            //免费礼包(邮件、展示)
	battleAttrMap        map[propertytypes.BattlePropertyType]int64 //vip等级属性
}

func (t *VipTemplate) TemplateId() int {
	return t.Id
}

func (t *VipTemplate) GetBattleAttrMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrMap
}

func (t *VipTemplate) GetDisCountItemMap() map[int32]int32 {
	return t.disCountItemMap
}

func (t *VipTemplate) GetFreeGiftItemMap() map[int32]int32 {
	return t.freeGiftItemMap
}

func (t *VipTemplate) GetEmailDisCountItemMap() map[int32]int32 {
	return t.emailDiscountItemMap
}

func (t *VipTemplate) GetEmailFreeGiftItemMap() map[int32]int32 {
	return t.emailFreeGiftItemMap
}

func (t *VipTemplate) GetNextTemplate() *VipTemplate {
	return t.nextTemp
}

func (t *VipTemplate) PatchAfterCheck() {
	if t.GiftSilver > 0 {
		t.emailDiscountItemMap[constanttypes.SilverItem] = int32(t.GiftSilver)
	}

	if t.FreeGiftSilver > 0 {
		t.emailFreeGiftItemMap[constanttypes.SilverItem] = int32(t.FreeGiftSilver)
	}
}
func (t *VipTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//属性
	t.battleAttrMap = make(map[propertytypes.BattlePropertyType]int64)
	if t.Hp > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	}
	if t.Attack > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	}
	if t.Defence > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	}

	//折扣物品
	t.disCountItemMap = make(map[int32]int32)
	t.emailDiscountItemMap = make(map[int32]int32)
	rewItemIdList, err := utils.SplitAsIntArray(t.GiftId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GiftId)
		return template.NewTemplateFieldError("GiftId", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(t.GiftCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.GiftCount)
		return template.NewTemplateFieldError("GiftCount", err)
	}
	if len(rewItemIdList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.GiftId, t.GiftCount)
		return template.NewTemplateFieldError("GiftId or GiftCount", err)
	}
	if len(rewItemIdList) > 0 {
		//组合数据
		for index, itemId := range rewItemIdList {
			_, ok := t.disCountItemMap[itemId]
			if ok {
				t.disCountItemMap[itemId] += rewItemCountList[index]
			} else {
				t.disCountItemMap[itemId] = rewItemCountList[index]
			}

			_, ok = t.emailDiscountItemMap[itemId]
			if ok {
				t.emailDiscountItemMap[itemId] += rewItemCountList[index]
			} else {
				t.emailDiscountItemMap[itemId] = rewItemCountList[index]
			}
		}
	}

	//免费礼包物品
	t.freeGiftItemMap = make(map[int32]int32)
	t.emailFreeGiftItemMap = make(map[int32]int32)
	freeItemIdList, err := utils.SplitAsIntArray(t.FreeGiftId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FreeGiftId)
		return template.NewTemplateFieldError("FreeGiftId", err)
	}
	freeItemCountList, err := utils.SplitAsIntArray(t.FreeGiftCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FreeGiftCount)
		return template.NewTemplateFieldError("FreeGiftCount", err)
	}
	if len(freeItemIdList) != len(freeItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.FreeGiftId, t.FreeGiftCount)
		return template.NewTemplateFieldError("FreeGiftId or FreeGiftCount", err)
	}
	if len(freeItemIdList) > 0 {
		//组合数据
		for index, itemId := range freeItemIdList {
			_, ok := t.freeGiftItemMap[itemId]
			if ok {
				t.freeGiftItemMap[itemId] += freeItemCountList[index]
			} else {
				t.freeGiftItemMap[itemId] = freeItemCountList[index]
			}

			_, ok = t.emailFreeGiftItemMap[itemId]
			if ok {
				t.emailFreeGiftItemMap[itemId] += freeItemCountList[index]
			} else {
				t.emailFreeGiftItemMap[itemId] = freeItemCountList[index]
			}
		}
	}

	return nil
}

func (t *VipTemplate) Check() (err error) {
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
		to := template.GetTemplateService().Get(int(t.NextId), (*VipTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextTemp = to.(*VipTemplate)
	}

	//验证 等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}
	//验证 星级
	err = validator.MinValidate(float64(t.Star), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Star)
		err = template.NewTemplateFieldError("Star", err)
		return
	}
	//验证 升级条件
	err = validator.MinValidate(float64(t.NeedValue), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedValue)
		err = template.NewTemplateFieldError("NeedValue", err)
		return
	}

	//验证 原价
	err = validator.MinValidate(float64(t.CostPrice), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.CostPrice)
		err = template.NewTemplateFieldError("CostPrice", err)
		return
	}

	//验证 现价
	err = validator.MinValidate(float64(t.Price), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Price)
		err = template.NewTemplateFieldError("Price", err)
		return
	}

	//验证 礼包物品
	for itemId, num := range t.disCountItemMap {
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("GiftId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			err = template.NewTemplateFieldError("GiftCount", err)
			return
		}
	}

	//验证 礼包银两
	err = validator.MinValidate(float64(t.GiftSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GiftSilver)
		err = template.NewTemplateFieldError("GiftSilver", err)
		return
	}

	//验证 免费礼包物品
	for itemId, num := range t.freeGiftItemMap {
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("FreeGiftId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			err = template.NewTemplateFieldError("FreeGiftCount", err)
			return
		}
	}

	//验证 免费礼包银两
	err = validator.MinValidate(float64(t.FreeGiftSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FreeGiftSilver)
		err = template.NewTemplateFieldError("FreeGiftSilver", err)
		return
	}

	//验证 手续
	err = validator.RangeValidate(float64(t.Shouxu), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Shouxu)
		err = template.NewTemplateFieldError("Shouxu", err)
		return
	}

	return nil
}

func (t *VipTemplate) FileName() string {
	return "tb_vip_info.json"
}

func init() {
	template.Register((*VipTemplate)(nil))
}
