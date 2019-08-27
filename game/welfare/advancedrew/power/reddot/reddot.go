package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	advancedrewpowertypes "fgame/fgame/game/welfare/advancedrew/power/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypePower, reddot.HandlerFunc(handleRedDotAdvancedRewPower))
}

//进阶战力奖励红点
func handleRedDotAdvancedRewPower(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedrewpowertypes.AdvancedPowerInfo)
	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needPower := int64(temp.Value2)
		if !info.IsCanReceiveRewards(needPower) {
			continue
		}

		isNotice = true
		return
	}

	return
}
