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
	investnewsevendaytemplate "fgame/fgame/game/welfare/invest/new_sevenday/template"
	investnewsevendaytypes "fgame/fgame/game/welfare/invest/new_sevenday/types"
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
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewServenDay, welfare.ReceiveHandlerFunc(investNewSevenDayReceive))
}

func investNewSevenDayReceive(pl player.Player, rewId int32) (err error) {
	playerId := pl.GetId()
	typ := welfaretypes.OpenActivityTypeInvest
	subType := welfaretypes.OpenActivityInvestSubTypeNewServenDay

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
	investType := openTemp.Value1

	// 校验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	info := obj.GetActivityData().(*investnewsevendaytypes.NewInvestDayInfo)

	// 判断是否购买
	if !info.IsAlreadyBuy(investType) {
		log.WithFields(
			log.Fields{
				"playerId":   playerId,
				"groupId":    groupId,
				"investType": investType,
			}).Warn("welfare: 新七日投资领取奖励请求，未购买该档次的新七日投资")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotBuyInvest)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*investnewsevendaytemplate.GroupTemplateNewInvestDay)
	maxRewards := groupTemp.GetInvestDayMaxRewardsLevel(investType)

	// 判断是否能领取的奖励
	if !info.IsCanReceive(investType, maxRewards) {
		log.WithFields(
			log.Fields{
				"playerId":   playerId,
				"groupId":    groupId,
				"investType": investType,
			}).Warn("welfare: 已经领取过该奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	maxExclude, _ := info.GetReceiveByType(investType)
	curDay, _ := info.GetCurDay(investType)

	showRewItemMap := make(map[int32]int32)
	rewItemMap := make(map[int32]int32)
	totalRewData := propertytypes.CreateRewData(0, 0, 0, 0, 0)
	tempList := groupTemp.GetInvestDayRewTempList(investType, maxExclude, curDay)
	for _, temp := range tempList {
		//资源
		totalRewData.RewBindGold += temp.RewGoldBind
		totalRewData.RewGold += temp.RewGold
		totalRewData.RewSilver += temp.RewSilver

		//物品
		for itemId, num := range temp.GetRewItemMap() {
			_, ok := rewItemMap[itemId]
			if ok {
				rewItemMap[itemId] += num
			} else {
				rewItemMap[itemId] = num
			}
		}

		//前端展示
		for itemId, num := range temp.GetEmailRewItemMap() {
			_, ok := showRewItemMap[itemId]
			if ok {
				showRewItemMap[itemId] += num
			} else {
				showRewItemMap[itemId] = num
			}
		}
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//背包空间
	firsOpenTemp := groupTemp.GetFirstOpenTemp()
	newItemDataList := welfarelogic.ConvertToItemData(rewItemMap, firsOpenTemp.GetExpireType(), firsOpenTemp.GetExpireTime())
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:新七日投资领取奖励请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
	itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
	flag := inventoryManager.BatchAddOfItemLevel(newItemDataList, itemGetReason, itemGetReasonText)
	if !flag {
		panic("welfare:new seven day invest rewards add item should be ok")
	}

	reasonGold := commonlog.GoldLogReasonOpenActivityRew
	reasonSilver := commonlog.SilverLogReasonOpenActivityRew
	reasonLevel := commonlog.LevelLogReasonOpenActivityRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), typ, subType)
	reasonSilverText := fmt.Sprintf(reasonSilver.String(), typ, subType)
	reasonLevelText := fmt.Sprintf(reasonLevel.String(), typ, subType)
	flag = propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSilverText, reasonLevel, reasonLevelText)
	if !flag {
		panic("welfare:new seven day invest rewards add RewData should be ok")
	}

	// 推送改变
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	// 刷新数据
	info.UpdateNewSevenDayInvest(investType, curDay)
	welfareManager.UpdateObj(obj)

	var record []int32
	record = append(record, curDay)

	scOpenActivityReceiveRew := pbutil.BuildSCOpenActivityReceiveRew(rewId, openTemp.Group, totalRewData, rewItemMap, record)
	pl.SendMsg(scOpenActivityReceiveRew)

	return
}
