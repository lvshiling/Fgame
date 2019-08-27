package handler

import (
	"fgame/fgame/common/lang"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	discountyunyintypes "fgame/fgame/game/welfare/discount/yunyin/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeYunYin, welfare.ReceiveHandlerFunc(yunYinReceiveReward))
}

func yunYinReceiveReward(pl player.Player, rewId int32) (err error) {
	playerId := pl.GetId()
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeYunYin

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
	goldNum := openTemp.Value1

	// 检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	info := obj.GetActivityData().(*discountyunyintypes.YunYinInfo)

	// 判断是否能领取该奖励
	if !info.IsCanReceive(goldNum) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"groupId":  groupId,
				"goldNum":  goldNum,
			}).Warn("welfare: 领取云隐商店奖励请求, 不能领取该奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityTongTianTaNotReceive)
		return
	}

	// 判断是否已经领取了该奖励
	if info.IsAlreadyReceive(goldNum) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"groupId":  groupId,
				"goldNum":  goldNum,
			}).Warn("welfare: 领取云隐商店奖励请求, 已经领取该奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityTongTianTaAlreadyReceive)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	// 推送改变
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	// 刷新数据
	info.AddReceiveRecord(goldNum)
	welfareManager.UpdateObj(obj)

	scOpenActivityReceiveRew := pbutil.BuildSCOpenActivityReceiveRew(rewId, openTemp.Group, totalRewData, rewItemMap, info.ReceiveRecord)
	pl.SendMsg(scOpenActivityReceiveRew)
	return
}
