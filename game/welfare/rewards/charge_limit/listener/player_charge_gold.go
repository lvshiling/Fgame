package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	chargelimitlogic "fgame/fgame/game/welfare/rewards/charge_limit/logic"
	rewardschargelimittypes "fgame/fgame/game/welfare/rewards/charge_limit/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家充值元宝
func playerChargeRewardsLimit(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}
	typ := welfaretypes.OpenActivityTypeRewards
	subType := welfaretypes.OpenActivityRewardsSubTypeChargeLimit

	//充值领奖（每多少领奖）(全服次数)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*rewardschargelimittypes.ChargeRewLimitInfo)
		info.GoldNum += goldNum
		info.AddLeftNum(goldNum)

		chargelimitlogic.CheckRewardsMail(obj)
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeRewardsLimit))
}
