package handler

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	chargesingleallmultipletemplate "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/template"
	chargesingleallmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeAllRew, welfare.ReceiveHandlerFunc(receviceHandler))
}

func receviceHandler(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeCycleCharge
	subType := welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeAllRew

	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取每日单笔充值奖励请求，模板不存在")
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
			}).Warn("welfare:领取每日单笔充值奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*chargesingleallmultipletypes.CycleChargeSingleAllMultipleInfo)

	cycDay := openTemp.Value1
	if cycDay != info.CycleDay {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"rewId":     rewId,
				"cycDay":    cycDay,
				"curCycDay": info.CycleDay,
			}).Warn("welfare:领取每日单笔充值奖励请求,充值日类型错误")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityCycleTypeNotEqual)
		return
	}
	needGoldNum := openTemp.Value2
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取每日单笔充值奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	groupTemp := groupInterface.(*chargesingleallmultipletemplate.GroupTemplateCycleSingleAllMultiple)

	remainTimesMap := info.GetCanRewRecord()
	useTimes, rewTimes, flag := groupTemp.GetSingleGoldCanRewRecordMap(info.CycleDay, needGoldNum, remainTimesMap)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"needGoldNum": needGoldNum,
			}).Warn("welfare:领取每日单笔充值奖励请求，不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}
	// canRewRecordMap := groupTemp.GetCanRewRecordMap(info.CycleDay, needGoldNum, remainTimesMap)
	// num := canRewRecordMap[needGoldNum]
	// if num <= 0 {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":    pl.GetId(),
	// 			"needGoldNum": needGoldNum,
	// 		}).Warn("welfare:领取每日单笔充值奖励请求，不满足领取条件")
	// 	playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
	// 	return
	// }

	// useTimesMap := make(map[int32]int32)
	// for goldNum, times := range remainTimesMap {
	// 	for _, temp := range descTempList {
	// 		needGold := temp.Value2
	// 		if goldNum < needGold {
	// 			continue
	// 		}
	// 		num := goldNum / needGold
	// 		canRewRecordMap[needGold] += num * times
	// 	}
	// }

	// descTempList := groupTemp.GetCurDayTempDescList(info.CycleDay)
	// for _, temp := range descTempList {
	// 	canRewGold := temp.Value2
	// 	if canRewGold != needGoldNum {
	// 		continue
	// 	}

	// 	//领取条件
	// 	if !info.IsCanReceiveRewards(needGoldNum) {
	// 		log.WithFields(
	// 			log.Fields{
	// 				"playerId":    pl.GetId(),
	// 				"needGoldNum": needGoldNum,
	// 			}).Warn("welfare:领取每日单笔充值奖励请求，不满足领取条件")
	// 		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
	// 		return
	// 	}
	// }

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//openTemp.GetEmailRewItemDataListWithRatio(rewTimes)
	itemDataList := welfarelogic.ConvertToItemDataWithWelfareData(openTemp.GetRewItemDataListWithRatio(rewTimes), openTemp.GetExpireType(), openTemp.GetExpireTime())

	if len(itemDataList) > 0 {
		//背包空间
		if !inventoryManager.HasEnoughSlotsOfItemLevel(itemDataList) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("welfare:领取活动奖励请求，背包空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
		flag := inventoryManager.BatchAddOfItemLevel(itemDataList, itemGetReason, itemGetReasonText)
		if !flag {
			panic("welfare:添加物品应该成功")
		}
	}

	reasonGold := commonlog.GoldLogReasonOpenActivityRew
	reasonSilver := commonlog.SilverLogReasonOpenActivityRew
	reasonLevel := commonlog.LevelLogReasonOpenActivityRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), typ, subType)
	reasonSilverText := fmt.Sprintf(reasonSilver.String(), typ, subType)
	reasonLevelText := fmt.Sprintf(reasonLevel.String(), typ, subType)

	rewSilver := openTemp.RewSilver * rewTimes
	rewBindGold := openTemp.RewGoldBind * rewTimes
	rewGold := openTemp.RewGold * rewTimes
	rewExp := int32(0)
	rewExpPoint := int32(0)
	totalRewData := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)
	flag = propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSilverText, reasonLevel, reasonLevelText)
	if !flag {
		panic("welfare:welfare rewards add RewData should be ok")
	}
	totalItemMap := openTemp.GetEmailRewItemMapWithRatio(rewTimes)

	info.Receive(useTimes)
	welfareManager.UpdateObj(obj)
	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	canRewRecord := groupTemp.GetCanRewRecordMap(info.CycleDay, info.GetCanRewRecord())
	scMsg := pbutil.BuildSCOpenActivityReceiveRewCycleSingleAllRewMultiple(rewId, groupId, totalRewData, totalItemMap, canRewRecord)
	pl.SendMsg(scMsg)
	return
}

func mergeMap(a, b map[int32]int32) map[int32]int32 {
	c := make(map[int32]int32)
	for key, _ := range b {
		c[key] = a[key] + b[key]
	}
	return c
}
