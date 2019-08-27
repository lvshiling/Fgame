package template

import (
	inventorytypes "fgame/fgame/game/inventory/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//超值套餐
type GroupTemplateDiscountTaoCan struct {
	*welfaretemplate.GroupTemplateBase
}

//套餐价格
func (gt *GroupTemplateDiscountTaoCan) GetNeedGold() int32 {
	return gt.GetFirstValue1()
}

//装备礼包索引
func (gt *GroupTemplateDiscountTaoCan) GetEquipGiftIndex() int32 {
	return gt.GetFirstValue2()
}

//非套餐折扣数
func (gt *GroupTemplateDiscountTaoCan) GetOptinalDiscountRate() int32 {
	return gt.GetFirstValue3()
}

//获取套餐额外奖励
func (gt *GroupTemplateDiscountTaoCan) GetRewItemMap() map[int32]int32 {
	return gt.GetFirstOpenTemp().GetRewItemMap()
}

//获取过期类型
func (gt *GroupTemplateDiscountTaoCan) GetExpireType() inventorytypes.NewItemLimitTimeType {
	return gt.GetFirstOpenTemp().GetExpireType()
}

//获取过期时间
func (gt *GroupTemplateDiscountTaoCan) GetExpireTime() int64 {
	return gt.GetFirstOpenTemp().GetExpireTime()
}

func CreateGroupTemplateDiscountTaoCan(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateDiscountTaoCan{}
	gt.GroupTemplateBase = base
	return gt
}
func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeTaoCan, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateDiscountTaoCan))
}
