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
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_FEEDBACK_DEVELOP_RECEIVE_TYPE), dispatch.HandlerFunc(handlerFeedbackDevelopReceive))
}

//处理领取金鸡
func handlerFeedbackDevelopReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理金鸡领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityFeedbackDevelopReceive)
	rewId := csMsg.GetRewId()

	err = receiveFeedbackDevelop(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理金鸡领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理金鸡领取奖励请求完成")

	return
}

//金鸡领取奖励请求逻辑
func receiveFeedbackDevelop(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeChargeDevelop
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取金鸡奖励请求，模板不存在")
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
			}).Warn("welfare:领取金鸡奖励请求，活动不存在")
		return
	}

	err = welfareManager.RefreshActivityDataByGroupId(openTemp.Group)
	if err != nil {
		return
	}

	info := obj.GetActivityData().(*feedbackchargedeveloptypes.FeedbackDevelopInfo)
	if !info.IsActivate {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  openTemp.Group,
			}).Warn("welfare:喂养金鸡请求，金鸡未激活")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityFeedbackDevelopNotActivate)
		return
	}

	if info.IsDead {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  openTemp.Group,
			}).Warn("welfare:喂养金鸡请求，金鸡已经死亡")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityFeedbackDevelopHadDead)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(openTemp.Group)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackchargedeveloptemplate.GroupTemplateDevelop)

	//领取条件
	switch feedbackchargedeveloptypes.FeedbackDevelopRewType(openTemp.Value1) {
	case feedbackchargedeveloptypes.FeedbackDevelopRewTypeSingleDay:
		{
			if info.IsFeed {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"groupId":  openTemp.Group,
					}).Warn("welfare:领取金鸡奖励请求，已经喂养金鸡")
				return
			}

			//喂养次数
			rewTimes := openTemp.Value2
			if info.FeedTimes != rewTimes {
				log.WithFields(
					log.Fields{
						"playerId":     pl.GetId(),
						"rewId":        rewId,
						"rewTimes":     rewTimes,
						"curFeedTimes": info.FeedTimes,
					}).Warn("welfare:领取金鸡奖励请求,喂养次数错误")
				playerlogic.SendSystemMessage(pl, lang.OpenActivityCycleTypeNotEqual)
				return
			}

			//喂养条件
			needCost := groupTemp.GetDevelopFeedCondition(info.FeedTimes)
			if info.TodayCostNum < needCost {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"needCost": needCost,
						"curCost":  info.TodayCostNum,
					}).Warn("welfare:喂养金鸡，没有饲料")
				playerlogic.SendSystemMessage(pl, lang.OpenActivityFeedbackDevelopNotFeedTimes)
				return
			}
		}
	case feedbackchargedeveloptypes.FeedbackDevelopRewTypeCountDay:
		{
			needFeedTimes := groupTemp.GetDevelopNeedTotalTimes()
			if info.FeedTimes != needFeedTimes {
				log.WithFields(
					log.Fields{
						"playerId":       pl.GetId(),
						"totalFeedTimes": info.FeedTimes,
						"needFeedTimes":  openTemp.Value2,
					}).Warn("welfare:喂养金鸡，喂养日不够")
				playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
				return
			}
		}
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	switch feedbackchargedeveloptypes.FeedbackDevelopRewType(openTemp.Value1) {
	case feedbackchargedeveloptypes.FeedbackDevelopRewTypeSingleDay:
		{
			info.IsFeed = true
			info.FeedTimes += 1
		}
	case feedbackchargedeveloptypes.FeedbackDevelopRewTypeCountDay:
		{
			info.IsReceiveRew = true
		}
	}
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityFeedbackDevelopReceive(totalRewData, rewItemMap)
	pl.SendMsg(scMsg)
	return
}
