package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargearenapvpassisttypesreturn "fgame/fgame/game/welfare/alliance/charge_arenapvp_assist_return/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeWuLian, welfare.InfoGetHandlerFunc(handlerChargeArenapvpAssistReturnGetInfo))
}

func handlerChargeArenapvpAssistReturnGetInfo(pl player.Player, groupId int32) (err error) {
	tpl := pl
	err = getInfo(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理擂台助力-消费助力获取信息请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理擂台助力-消费助力获取信息请求完成")

	return
}

func getInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeAlliance
	subType := welfaretypes.OpenActivityAllianceSubTypeWuLian
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
			}).Warn("welfare:领取擂台助力-消费助力奖励请求，活动不存在")
		return
	}

	info := obj.GetActivityData().(*feedbackchargearenapvpassisttypesreturn.FeedbackChargeArenapvpAssistReturnInfo)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scMsg := pbutil.BuildSCOpenActivityFeedbackChargeArenapvpAssistReturnInfo(groupId, startTime, endTime, int32(info.CostNum))
	pl.SendMsg(scMsg)
	return
}
