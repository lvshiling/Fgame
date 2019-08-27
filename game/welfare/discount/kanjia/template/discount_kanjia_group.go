package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//充值翻倍
type GroupTemplateDiscountKanJia struct {
	*welfaretemplate.GroupTemplateBase
}

// 打折礼包-赠送条件
func (gt *GroupTemplateDiscountKanJia) GetBargainRewTimesNeedGold() int32 {
	return gt.GetFirstValue1()
}

func CreateGroupTemplateDiscountKanJia(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateDiscountKanJia{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeKanJia, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateDiscountKanJia))
}
