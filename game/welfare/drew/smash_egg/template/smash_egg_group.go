package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//砸金蛋
type GroupTemplateSmashEgg struct {
	*welfaretemplate.GroupTemplateBase
}

//砸金蛋一批的数量
func (gt *GroupTemplateSmashEgg) GetSmashEggBatchCount() int32 {
	return gt.GetFirstValue1()
}

//砸金蛋元宝价格
func (gt *GroupTemplateSmashEgg) GetNeedGold() int32 {
	return gt.GetFirstValue2()
}

//砸金蛋绑元价格
func (gt *GroupTemplateSmashEgg) GetNeedBindGold() int32 {
	return gt.GetFirstValue3()
}

//砸金蛋银两价格
func (gt *GroupTemplateSmashEgg) GetNeedSilver() int32 {
	return gt.GetFirstValue4()
}

func CreateGroupTemplateSmashEgg(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateSmashEgg{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmashEgg, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateSmashEgg))
}
