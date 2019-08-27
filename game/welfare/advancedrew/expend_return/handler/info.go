package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	advancedrewexpendreturntypes "fgame/fgame/game/welfare/advancedrew/expend_return/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeExpendReturn, welfare.InfoGetHandlerFunc(expendReturnInfo))
}

//升阶消耗返还信息
func expendReturnInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityDataByGroupId(groupId)
	if err != nil {
		return
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	var record []int32
	danNum := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*advancedrewexpendreturntypes.AdvancedExpendReturnInfo)
		record = info.RewRecord
		danNum = info.DanNum
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoAdvancedExpendReturn(groupId, startTime, endTime, record, danNum)
	pl.SendMsg(scMsg)
	return
}
