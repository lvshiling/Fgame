package handler

import (
	"fgame/fgame/common/lang"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
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
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCharge, welfare.ReceiveHandlerFunc(handlerChargeArenapvpAssistReceive))
}

func handlerChargeArenapvpAssistReceive(pl player.Player, rewId int32) (err error) {
	tpl := pl

	err = receiveChargeArenapvpAssistFeedback(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理擂台助力领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理擂台助力领取奖励请求完成")

	return

}

func receiveChargeArenapvpAssistFeedback(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeCharge
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取擂台助力奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, openTemp.Group)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(openTemp.Group)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  openTemp.Group,
			}).Warn("welfare:领取擂台助力奖励请求，活动不存在")
		return
	}

	info := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)
	needGoldNum := openTemp.Value1

	if !info.IsCanReceiveRewards(needGoldNum) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
				"goldNum":  needGoldNum,
				"group":    openTemp.Group,
			}).Warn("welfare:领取擂台助力奖励请求，不满足领取条件(充值不足或已领取)")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新数据库信息
	info.AddRecord(needGoldNum)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityReceiveRew(rewId, openTemp.Group, totalRewData, rewItemMap, info.RewRecord)
	pl.SendMsg(scMsg)

	return
}
