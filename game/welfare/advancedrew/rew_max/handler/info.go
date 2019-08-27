package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	advancedrewrewmaxtypes "fgame/fgame/game/welfare/advancedrew/rew_max/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewMax, welfare.InfoGetHandlerFunc(getAdvancedRewMaxInfo))
}

//获取升阶奖励请求逻辑
func getAdvancedRewMaxInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.RefreshActivityDataByGroupId(groupId)
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeRewMax

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	var record []int32
	chargeNum := int32(0)
	initAdvanced := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*advancedrewrewmaxtypes.AdvancedRewMaxInfo)
		chargeNum = info.PeriodChargeNum
		initAdvanced = info.InitAdvancedNum
		record = info.RewRecord
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoAdvancedRewMax(groupId, startTime, endTime, record, chargeNum, initAdvanced)
	pl.SendMsg(scMsg)
	return
}
