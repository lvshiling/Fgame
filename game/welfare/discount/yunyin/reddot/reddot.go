package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	discountyunyintemplate "fgame/fgame/game/welfare/discount/yunyin/template"
	discountyunyintypes "fgame/fgame/game/welfare/discount/yunyin/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeYunYin, reddot.HandlerFunc(handleReddotAboutYunYinShop))
}

func handleReddotAboutYunYinShop(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	if !welfarelogic.IsOnActivityTime(groupId) {
		return
	}
	info := obj.GetActivityData().(*discountyunyintypes.YunYinInfo)

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	// 判断是否能领取
	yunYinTemp := groupInterface.(*discountyunyintemplate.GroupTemplateDiscountYunYinShop)
	costList := yunYinTemp.GetCanReceiveRewardList(info.GoldNum)
	for _, cost := range costList {
		if !info.IsCanReceive(cost) {
			continue
		}
		if info.IsAlreadyReceive(cost){
			continue
		}
		isNotice = true
		return
	}

	return
}
