package handler

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	systemlingyutypes "fgame/fgame/game/welfare/system/lingyu/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeSystemActivate, welfaretypes.OpenActivitySystemActivateSubTypeLingYu, welfare.InfoGetHandlerFunc(handlerActivateLingyuInfo))
}

//领域系统激活信息请求
func handlerActivateLingyuInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	var record []int32
	startTime, endTime := int64(0), int64(0)
	isActivate := false
	maxSingle := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*systemlingyutypes.SystemLingYuInfo)
		isActivate = info.IsActivate
		maxSingle = info.MaxSingleChargeGold
		startTime = info.StartTime

		openTempMap := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
		if len(openTempMap) != 1 {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Error("welfare:领域活动激活模板应该只有一条")
			return
		}
		for _, openTemp := range openTempMap {
			continuedTime := int64(openTemp.Value2) * int64(common.DAY)
			endTime = info.StartTime + continuedTime
		}
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoSystemActivate(groupId, startTime, endTime, record, maxSingle, isActivate)
	pl.SendMsg(scMsg)
	return
}
