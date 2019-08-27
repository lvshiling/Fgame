package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//登录奖励
type GroupTemplateHallLogin struct {
	*welfaretemplate.GroupTemplateBase //
}

//登录奖励最大天数
func (gt *GroupTemplateHallLogin) GetWelfareLoginMaxDay() int32 {
	maxDay := gt.GetMaxValue1()
	return maxDay
}

func CreateGroupTemplateHallLogin(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	g := &GroupTemplateHallLogin{}
	g.GroupTemplateBase = base
	return g
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeLogin, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateHallLogin))
}
