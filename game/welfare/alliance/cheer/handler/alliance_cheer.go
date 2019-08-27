package info

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	alliancecheertypes "fgame/fgame/game/welfare/alliance/cheer/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeAlliance, welfare.InfoGetHandlerFunc(handlerAllianceCheerInfo))
}

//城战助威信息请求
func handlerAllianceCheerInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	var record []int32
	poolNum := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*alliancecheertypes.AllianceCheerInfo)
		poolNum = info.CheerGoldNum
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoAllianceCheerInfo(groupId, startTime, endTime, record, poolNum)
	pl.SendMsg(scMsg)
	return
}
