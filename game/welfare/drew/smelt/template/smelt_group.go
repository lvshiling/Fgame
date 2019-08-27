package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type GroupTemplateSmelt struct {
	*welfaretemplate.GroupTemplateBase
}

func CreateGroupTemplateSmelt(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateSmelt{}
	gt.GroupTemplateBase = base
	return gt
}

func (gt *GroupTemplateSmelt) GetItemId() int32 {
	return gt.GetFirstValue1()
}

func (gt *GroupTemplateSmelt) GetNeedItemNum() int32 {
	return gt.GetFirstValue2()
}

func (gt *GroupTemplateSmelt) Init() (err error) {
	if len(gt.GetOpenTempMap()) != 1 {
		panic("welfare:炼金炉-冶炼应该只有一条配置")
	}
	return
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmelt, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateSmelt))
}
