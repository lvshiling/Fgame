package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	advancedrewrewmaxtemplate "fgame/fgame/game/welfare/advancedrew/rew_max/template"
	advancedrewrewmaxtypes "fgame/fgame/game/welfare/advancedrew/rew_max/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewMax, reddot.HandlerFunc(handleRedDotAdvancedRewMax))
}

func handleRedDotAdvancedRewMax(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewrewmaxtemplate.GroupTemplateRewMax)
	info := obj.GetActivityData().(*advancedrewrewmaxtypes.AdvancedRewMaxInfo)
	if info.RewType != groupTemp.GetAdvancedType() {
		return
	}

	//最低至初始阶数的档次
	minRewAdvanced := int32(0)
	for _, temp := range groupTemp.GetRewTempDescList() {
		if temp.Value2 <= info.InitAdvancedNum {
			minRewAdvanced = temp.Value2
			break
		}
	}

	for _, temp := range groupTemp.GetRewTempDescList() {
		needAdvancedNum := temp.Value2
		needChargeNum := temp.Value3

		// 初始条件最近的档次
		if needAdvancedNum < minRewAdvanced {
			break
		}

		if info.IsCanReceiveRewards(needAdvancedNum, needChargeNum) {
			isNotice = true
			return
		}

	}

	return
}
