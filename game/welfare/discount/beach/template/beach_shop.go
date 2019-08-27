package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type GroupTemplateDiscountBeachShop struct {
	*welfaretemplate.GroupTemplateBase
}

func (t *GroupTemplateDiscountBeachShop) GetAvtiviteItemMap() map[int32]int32 {
	itemMap := make(map[int32]int32)

	itemId := t.GetFirstValue1()
	itemCount := t.GetFirstValue2()
	itemMap[itemId] = itemCount

	return itemMap
}

func CreateGroupTemplateDiscountBeachShop(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	t := &GroupTemplateDiscountBeachShop{}
	t.GroupTemplateBase = base
	return t

}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeBeach, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateDiscountBeachShop))
}
