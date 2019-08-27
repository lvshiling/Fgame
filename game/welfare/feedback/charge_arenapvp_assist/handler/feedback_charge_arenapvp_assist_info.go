package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargetypes "fgame/fgame/game/welfare/feedback/charge/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCharge, welfare.InfoGetHandlerFunc(handlerChargeArenapvpAssistGetInfo))
}

func handlerChargeArenapvpAssistGetInfo(pl player.Player, groupId int32) (err error) {
	tpl := pl

	err = getFeedbackChargeArenapvpAssistInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理擂台助力获取信息请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理擂台助力获取信息请求完成")

	return
}

func getFeedbackChargeArenapvpAssistInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeCharge
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"groupId":  groupId,
		}).Warn("welfare:参数无效,活动时间模板不存在")
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:擂台助力信息请求，活动不存在")
		return
	}

	info := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scMsg := pbutil.BuildSCOpenActivityFeedbackChargeArenapvpAssistInfo(groupId, startTime, endTime, info.RewRecord, info.GoldNum)
	pl.SendMsg(scMsg)

	return
}
