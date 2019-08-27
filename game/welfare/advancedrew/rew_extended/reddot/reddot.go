package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	advancedrewrewextendedtypes "fgame/fgame/game/welfare/advancedrew/rew_extended/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewExtended, reddot.HandlerFunc(handleRedDotAdvancedRewExtended))
}

//进阶奖励红点
func handleRedDotAdvancedRewExtended(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	info := obj.GetActivityData().(*advancedrewrewextendedtypes.AdvancedRewExtendedInfo)
	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needAdvancedNum := temp.Value2
		if !info.IsCanReceiveRewards(needAdvancedNum) {
			continue
		}
		isNotice = true
		return
	}

	return
}
