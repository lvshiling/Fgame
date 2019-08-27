package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/pbutil"
	"fgame/fgame/game/reddot/reddot"
	hallupleveltypes "fgame/fgame/game/welfare/hall/uplevel/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeUpLevel, reddot.HandlerFunc(handleRedDotWelfareUplevel))
}

//福利大厅升级礼包红点
func handleRedDotWelfareUplevel(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	//升级礼包
	info := obj.GetActivityData().(*hallupleveltypes.WelfareUplevelInfo)
	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needLevel := temp.Value1
		if needLevel > pl.GetLevel() {
			continue
		}

		if !info.IsReceive(needLevel) {
			isNotice = true
		}
	}

	scMsg := pbutil.BuildSCActivityDataNotice(groupId, info.RewRecord)
	pl.SendMsg(scMsg)
	return
}
