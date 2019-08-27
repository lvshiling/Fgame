package template

import (
	"fgame/fgame/game/common/common"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//在线抽奖
type GroupTemplateWelfareOnline struct {
	*welfaretemplate.GroupTemplateBase
}

func CreateGroupTemplateWelfareOnline(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	g := &GroupTemplateWelfareOnline{}
	g.GroupTemplateBase = base
	return g
}

//在线抽奖次数
func (gt *GroupTemplateWelfareOnline) GetOpenActivityWelfareOnlineDrewTimes(onlineTime int64) int32 {
	drewTimes := int32(0)
	for _, temp := range gt.GetOpenTempMap() {
		conditionTime := int64(temp.Value1) * int64(common.SECOND)
		if conditionTime <= onlineTime {
			drewTimes += temp.Value2
		}
	}

	return drewTimes
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeOnline, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateWelfareOnline))
}
