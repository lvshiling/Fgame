package player

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardschargelimittypes "fgame/fgame/game/welfare/rewards/charge_limit/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 充值返利(全服次数)
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeChargeLimit, playerwelfare.ActivityObjInfoInitFunc(rewardsChargeLimitInitInfo))
}

func rewardsChargeLimitInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*rewardschargelimittypes.ChargeRewLimitInfo)
	info.GoldNum = 0
	info.TotalReceiveTimes = 0
	info.LeftConvertNumMap = map[int32]int32{}
	info.ReceiveTimesMap = map[int32]int32{}

	groupId := obj.GetGroupId()
	openTempMap := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, openTemp := range openTempMap {
		info.LeftConvertNumMap[openTemp.Value1] = 0
		info.ReceiveTimesMap[openTemp.Value1] = 0
	}
}
