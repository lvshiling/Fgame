package handler

import (
	"fgame/fgame/common/lang"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	goaltypes "fgame/fgame/game/welfare/huhu/goal/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeGoal, welfare.ReceiveHandlerFunc(receiveGoal))
}

//活动目标领取奖励请求逻辑
func receiveGoal(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeHuHu
	subType := welfaretypes.OpenActivitySpecialSubTypeGoal
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取活动目标奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
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
			}).Warn("welfare:领取活动目标奖励请求，活动不存在")
		return
	}

	//领取条件
	info := obj.GetActivityData().(*goaltypes.GoalInfo)
	rewGoalCount := openTemp.Value2
	if !info.IsCanReceiveRewards(rewGoalCount) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"rewId":        rewId,
				"rewGoalCount": rewGoalCount,
			}).Warn("welfare:领取活动目标奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddRecord(rewGoalCount)
	welfareManager.UpdateObj(obj)

	//同步资源
	inventorylogic.SnapInventoryChanged(pl)

	//同步属性
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCOpenActivityReceiveRew(rewId, openTemp.Group, totalRewData, rewItemMap, info.GetRewRecord())
	pl.SendMsg(scMsg)
	return
}
