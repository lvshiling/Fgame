package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	feedbackchargedeveloptemplate "fgame/fgame/game/welfare/feedback/charge_develop/template"
	feedbackchargedeveloptypes "fgame/fgame/game/welfare/feedback/charge_develop/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_FEEDBACK_DEVELOP_REVIVE_TYPE), dispatch.HandlerFunc(handlerFeedbackDevelopRevive))
}

//处理复活金鸡
func handlerFeedbackDevelopRevive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理复活金鸡请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityFeedbackDevelopRevive)
	groupId := csMsg.GetGroupId()

	err = developRevive(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理复活金鸡请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理复活金鸡请求完成")

	return
}

//复活金鸡请求逻辑
func developRevive(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeChargeDevelop

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:复活金鸡请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:复活金鸡请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*feedbackchargedeveloptypes.FeedbackDevelopInfo)
	if !info.IsCanRevive() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:复活金鸡请求，金鸡未死亡")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityFeedbackDevelopNotDead)
		return
	}

	groupTemp := groupInterface.(*feedbackchargedeveloptemplate.GroupTemplateDevelop)
	condition := groupTemp.GetReviveNeedGold()
	if info.ActivateChargeNum < condition {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:复活金鸡请求，激活条件不够")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityFeedbackDevelopCanNotActivate)
		return
	}
	//不需要消耗

	//复活条件
	// needGold := groupTemp.GetFeedbackDevelopReviveNeedGold()
	// propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	// if !propertyManager.HasEnoughGold(int64(needGold), false) {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"needGold": needGold,
	// 		}).Warn("welfare:复活金鸡请求，当前元宝不足")
	// 	playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
	// 	return
	// }

	// //消耗元宝
	// goldReason := commonlog.GoldLogReasonActivityCost
	// goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
	// flag := propertyManager.CostGold(int64(needGold), false, goldReason, goldReasonText)
	// if !flag {
	// 	panic("welfare:金鸡复活消耗应该成功")
	// }

	//复活金鸡
	info.IsActivate = true
	info.IsDead = false
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityFeedbackDevelopRevive(groupId, info.FeedTimes, info.IsDead)
	pl.SendMsg(scMsg)
	return
}
