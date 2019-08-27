package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	halllogintemplate "fgame/fgame/game/welfare/hall/login/template"
	halllogintypes "fgame/fgame/game/welfare/hall/login/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeLogin, reddot.HandlerFunc(handleRedDotWelfareLogin))
}

//福利大厅登录红点
func handleRedDotWelfareLogin(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	//登录奖励
	loginDay := welfarelogic.CountWelfareLoginDay(pl.GetCreateTime())
	loginInfo := obj.GetActivityData().(*halllogintypes.WelfareLoginInfo)
	if len(loginInfo.RewRecord) >= int(loginDay) {
		return
	}
	//最大登录奖励时间
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*halllogintemplate.GroupTemplateHallLogin)
	maxDay := groupTemp.GetWelfareLoginMaxDay()
	if len(loginInfo.RewRecord) >= int(maxDay) {
		return
	}

	isNotice = true
	return
}
