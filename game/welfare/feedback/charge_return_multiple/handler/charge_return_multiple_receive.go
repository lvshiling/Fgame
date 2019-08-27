package handler

import (
	"fgame/fgame/common/lang"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	feedbackchargereturnmultipletemplate "fgame/fgame/game/welfare/feedback/charge_return_multiple/template"
	feedbackchargereturnmultipletypes "fgame/fgame/game/welfare/feedback/charge_return_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple, welfare.ReceiveHandlerFunc(handlerChargeReturnMultipleReceive))
}

//循环充值领取奖励请求逻辑
func handlerChargeReturnMultipleReceive(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取循环充值奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	groupId := openTemp.Group

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
			}).Warn("welfare:领取循环充值奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*feedbackchargereturnmultipletypes.FeedbackChargeReturnMultipleInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取循环充值奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	groupTemp := groupInterface.(*feedbackchargereturnmultipletemplate.GroupTemplateChargeReturnMultiple)
	totalCnt := info.PeriodChargeNum / groupTemp.GetPerChargeNum()
	rewardLimitCnt := groupTemp.GetRewardLimitCnt()
	if rewardLimitCnt <= info.RewardCnt {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取循环充值奖励请求，已达领取次数上限")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityChargeReturnMultipleRewardLimit)
		return
	}

	if totalCnt <= info.RewardCnt {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取循环充值奖励请求，充值数不足")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityChargeReturnMultipleChargeNoEnough)
		return
	}

	leftCnt := totalCnt - info.RewardCnt
	if totalCnt > rewardLimitCnt {
		leftCnt = rewardLimitCnt - info.RewardCnt
	}

	totalRewData := propertytypes.CreateRewData(0, 0, 0, 0, 0)
	rewItemMap := make(map[int32]int32)
	addOk := false
	for i := int32(0); i < leftCnt; i++ {
		oneRewData, oneItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
		if !flag {
			break
		}
		info.AddRewardCnt(1)
		addOk = true
		for itemId, num := range oneItemMap {
			rewItemMap[itemId] += num
		}
		totalRewData.AddRewData(oneRewData)
	}
	if !addOk {
		return
	}

	//更新信息
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityReceiveRew(rewId, groupId, totalRewData, rewItemMap, nil)
	pl.SendMsg(scMsg)

	return
}
