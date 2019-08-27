package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	hallrealmtypes "fgame/fgame/game/welfare/hall/realm/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeRealm, reddot.HandlerFunc(handleRedDotOpenWelfareRealm))
}

//开服天劫塔冲刺红点
func handleRedDotOpenWelfareRealm(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	info := obj.GetActivityData().(*hallrealmtypes.WelfareRealmChallengeInfo)
	openTempMap := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, openTemp := range openTempMap {
		timesMax := openTemp.Value2
		floor := openTemp.Value1
		if !welfare.GetWelfareService().IsHadReceiveTimes(groupId, floor, timesMax, 1) {
			continue
		}

		if !info.IsCanReceiveRewards(floor) {
			continue
		}

		isNotice = true
	}

	return
}
