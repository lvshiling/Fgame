package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	hallzhuanshengtypes "fgame/fgame/game/welfare/hall/zhuansheng/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeZhaunSheng, welfare.InfoGetHandlerFunc(handlerZhuanShengInfo))
}

//转生冲刺信息请求
func handlerZhuanShengInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	var record []int32
	zhuanSheng := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*hallzhuanshengtypes.ZhuanShengInfo)
		record = info.RewRecord
		zhuanSheng = info.ZhuanSheng
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoZhuanSheng(groupId, startTime, endTime, record, zhuanSheng)
	pl.SendMsg(scMsg)
	return
}
