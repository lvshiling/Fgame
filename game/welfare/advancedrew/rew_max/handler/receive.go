package handler

import (
	"fgame/fgame/common/lang"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	advancedrewrewmaxtemplate "fgame/fgame/game/welfare/advancedrew/rew_max/template"
	advancedrewrewmaxtypes "fgame/fgame/game/welfare/advancedrew/rew_max/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewMax, welfare.ReceiveHandlerFunc(receiveAdvancedRewMax))
}

//升阶奖励领取奖励请求逻辑
func receiveAdvancedRewMax(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeRewMax
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取升阶奖励奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	groupId := openTemp.Group

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
				"groupId":  groupId,
			}).Warn("welfare:运营活动,不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取升阶奖励奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*advancedrewrewmaxtypes.AdvancedRewMaxInfo)

	advancedDay := welfaretypes.AdvancedType(openTemp.Value1)
	if advancedDay != info.RewType {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"rewId":           rewId,
				"advancedDay":     advancedDay,
				"curAdvancedType": info.RewType,
			}).Warn("welfare:领取升阶奖励奖励请求,升级奖励类型错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//领取条件
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取升阶奖励奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	groupTemp := groupInterface.(*advancedrewrewmaxtemplate.GroupTemplateRewMax)
	needAdvancedNum := openTemp.Value2
	needChargeNum := openTemp.Value3
	minRewAdvanced := int32(0)
	for _, temp := range groupTemp.GetRewTempDescList() {
		//最低至初始阶数的档次
		if temp.Value2 <= info.InitAdvancedNum {
			minRewAdvanced = temp.Value2
			break
		}
	}

	if needAdvancedNum < minRewAdvanced {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"minRewAdvanced":  minRewAdvanced,
				"needAdvancedNum": needAdvancedNum,
			}).Warn("welfare:领取升阶奖励奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	//领取条件
	if !info.IsCanReceiveRewards(needAdvancedNum, needChargeNum) {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"needAdvancedNum": needAdvancedNum,
				"rewId":           rewId,
			}).Warn("welfare:领取升阶奖励奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddRecord(needAdvancedNum)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityReceiveRew(rewId, groupId, totalRewData, rewItemMap, info.RewRecord)
	pl.SendMsg(scMsg)

	return
}
