package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	drewcommontypes "fgame/fgame/game/welfare/drew/common/types"
	drewcostdrewtemplate "fgame/fgame/game/welfare/drew/cost_drew/template"
	drewcostdrewtypes "fgame/fgame/game/welfare/drew/cost_drew/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_COST_DREW_ATTEND_TYPE), dispatch.HandlerFunc(handleAttendCostDrew))

}

//消费抽奖
func handleAttendCostDrew(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:消费抽奖")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityCostDrewAttend)
	groupId := csMsg.GetGroupId()
	drewInt := csMsg.GetDrewType()
	lastLogTime := csMsg.GetLastLogTime()
	isAutoBuy := csMsg.GetIsAutoBuy()
	dropLevel := csMsg.GetDropLevel()

	drewType := drewcommontypes.LuckyDrewAttendType(drewInt)
	if !drewType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:消费抽奖错误，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = attendCostDrew(tpl, groupId, dropLevel, drewType, lastLogTime, isAutoBuy)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("welfare:处理消费抽奖,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理消费抽奖完成")
	return nil

}

//消费抽奖逻辑：十连抽不支持三倍机制，需配置屏蔽掉
func attendCostDrew(pl player.Player, groupId, dropLevel int32, drewType drewcommontypes.LuckyDrewAttendType, lastLogTime int64, isAutoBuy bool) (err error) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeCostDrew

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:消费抽奖错误，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*drewcostdrewtypes.LuckyCostDrewInfo)
	luckTemp := welfaretemplate.GetWelfareTemplateService().GetLuckDrewTemplateByArg(groupId, dropLevel)
	if luckTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:消费抽奖错误，抽奖模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
	}
	ruleCycle := luckTemp.RewCount
	attendTimes := drewType.GetAttendNum()

	groupInteface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInteface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:消费抽奖错误，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	isOneCondition := false
	// 次数
	totalCostTimes := luckTemp.CostTimes * attendTimes
	if totalCostTimes > 0 {
		if info.LeftTimes < totalCostTimes {
			log.WithFields(
				log.Fields{
					"playerId":       pl.GetId(),
					"leftTimes":      info.LeftTimes,
					"totalCostTimes": totalCostTimes,
				}).Warn("welfare:消费抽奖错误，抽奖次数不足")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityNotEnoughDrewTimes)
			return
		}
		isOneCondition = true
	}

	//消耗物品
	shopIdMap := make(map[int32]int32)
	needItemMap := make(map[int32]int32)
	needGold := int64(0)
	needBindGold := int64(0)
	needSilver := int64(0)
	for itemId, num := range luckTemp.GetUseItemMap() {
		_, ok := needItemMap[itemId]
		if ok {
			needItemMap[itemId] += num * attendTimes
		} else {
			needItemMap[itemId] = num * attendTimes
		}
	}
	if len(needItemMap) > 0 {
		for needItemId, needNum := range needItemMap {
			totalNum := inventoryManager.NumOfItems(needItemId)
			if totalNum < needNum {
				if !isAutoBuy {
					log.WithFields(
						log.Fields{
							"playerId":    pl.GetId(),
							"needItemMap": needItemMap,
							"leftTimes":   info.LeftTimes,
						}).Warn("welfare:消费抽奖错误，物品不足")
					playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
					return
				}

				//自动购买
				if totalNum > 0 {
					needItemMap[needItemId] = totalNum
				} else {
					delete(needItemMap, needItemId)
				}
				needBuyNum := needNum - totalNum
				if needBuyNum > 0 {
					if !shop.GetShopService().ShopIsSellItem(needItemId) {
						log.WithFields(log.Fields{
							"playerId":  pl.GetId(),
							"isAutoBuy": isAutoBuy,
						}).Warn("chess:商铺没有该道具,无法自动购买")
						playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
						return
					}

					isEnoughBuyTimes, shopIdMap := shoplogic.MaxBuyTimesForPlayer(pl, needItemId, needBuyNum)
					if !isEnoughBuyTimes {
						log.WithFields(log.Fields{
							"playerId":  pl.GetId(),
							"isAutoBuy": isAutoBuy,
						}).Warn("chess:购买物品失败,抽奖失败")
						playerlogic.SendSystemMessage(pl, lang.ShopChessAutoBuyItemFail)
						return
					}

					shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
					needGold += shopNeedGold
					needBindGold += shopNeedBindGold
					needSilver += shopNeedSilver
				}
			}
		}

		isOneCondition = true
	}

	//是否足够元宝
	groupTemp := groupInteface.(*drewcostdrewtemplate.GroupTemplateCostDrew)
	needGold += groupTemp.GetCostDrewNeedGold() * int64(attendTimes)
	if needGold > 0 {
		if !propertyManager.HasEnoughGold(needGold, false) {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"leftTimes": info.LeftTimes,
					"needGold":  needGold,
				}).Warn("welfare:消费抽奖错误，元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}

		isOneCondition = true
	}

	if !isOneCondition {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"groupId":        groupId,
				"dropLevel":      dropLevel,
				"needGold":       needGold,
				"needItemMap":    len(needItemMap),
				"totalCostTimes": totalCostTimes,
				"leftTimes":      info.LeftTimes,
			}).Warn("welfare:消费抽奖错误，不满足抽奖条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityCostNotEnoughCondition)
		return
	}

	//是否足够银两
	if needSilver > 0 {
		flag := propertyManager.HasEnoughSilver(needSilver)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("chess:消费抽奖错误，银两不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//是否足够绑元
	if needBindGold > 0 {
		needCostBindGold := needBindGold + needGold
		flag := propertyManager.HasEnoughGold(needCostBindGold, true)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("chess:消费抽奖错误，绑元不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	if attendTimes > 0 {
		if inventoryManager.GetEmptySlots(inventorytypes.BagTypePrim) < attendTimes {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("welfare:充值抽奖错误，空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnoughSlot, fmt.Sprintf("%d", attendTimes))
			return
		}
	}

	isDrop := true
	isRandomResult := false
	// 不支持十连
	if attendTimes == 1 {
		// 是否规则
		if info.Ratio < 1 && info.AttendTimes > 0 && !info.IsRuleCD(ruleCycle) {
			needDouble := needGold * 2
			if info.LeftTimes < 1 && !propertyManager.HasEnoughGold(needDouble, false) {
				isDrop = false
				info.Ratio = luckTemp.RewTimes1
			} else {
				isRandomResult = true
			}
		}
	}

	if isRandomResult {
		switch luckTemp.GetRandomResultType() {
		case welfaretypes.DrewResultTypeDrop:
			{
				isDrop = true
			}
		case welfaretypes.DrewResultTypeRatioFirst:
			{
				isDrop = false
				info.Ratio = luckTemp.RewTimes1
			}
		case welfaretypes.DrewResultTypeRatioSecond:
			{
				isDrop = false
				info.Ratio = luckTemp.RewTimes2
			}
		}
	}

	// 第一次直接走掉落
	var totalItemList []*droptemplate.DropItemData
	var dropItemList []*droptemplate.DropItemData
	var extraItemList []*droptemplate.DropItemData
	var rewIndexList []int32
	if isDrop || info.AttendTimes == 0 {
		for times := int32(1); times <= attendTimes; times++ {
			rewIndex, itemData := countDropItem(pl, info, groupId, dropLevel)
			rewIndexList = append(rewIndexList, rewIndex)
			totalItemList = append(totalItemList, itemData)
			dropItemList = append(dropItemList, itemData)
		}

		info.CycleCount += attendTimes
	}

	for times := int32(1); times <= attendTimes; times++ {
		//每次额外奖励
		giveItemMap := luckTemp.GetGiveItemMap()
		if len(giveItemMap) > 0 {
			tempExtra := droptemplate.ConvertToItemDataList(giveItemMap, itemtypes.ItemBindTypeUnBind)
			extraItemList = append(extraItemList, tempExtra...)
			totalItemList = append(totalItemList, tempExtra...)
		}
	}

	var newItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(totalItemList) != 0 {
		newItemList, resMap = droplogic.SeperateItemDatas(totalItemList)
	}

	// 背包空间
	if len(newItemList) > 0 && !inventoryManager.HasEnoughSlotsOfItemLevel(newItemList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"len":      len(newItemList),
			}).Warn("welfare:消费抽奖错误，空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	if needGold > 0 {
		useGoldReason := commonlog.GoldLogReasonDrewUse
		useGoldReasonText := fmt.Sprintf(useGoldReason.String(), subType)
		flag := propertyManager.CostGold(needGold, false, useGoldReason, useGoldReasonText)
		if !flag {
			panic("welfare:幸运转盘消耗元宝应该成功")
		}
	}

	if len(needItemMap) > 0 {
		useItemReason := commonlog.InventoryLogReasonOpenActivityUse
		useItemReasonText := fmt.Sprintf(useItemReason.String(), typ, subType)
		flag := inventoryManager.BatchRemove(needItemMap, useItemReason, useItemReasonText)
		if !flag {
			panic("welfare:抽奖批量消耗物品应该成功")
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//增加掉落
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonOpenActivityRew
		silverReason := commonlog.SilverLogReasonOpenActivityRew
		levelReason := commonlog.LevelLogReasonOpenActivityRew
		goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
		silverReasonText := fmt.Sprintf(silverReason.String(), typ, subType)
		levelReasonText := fmt.Sprintf(levelReason.String(), typ, subType)
		err = droplogic.AddRes(pl, resMap, goldReason, goldReasonText, silverReason, silverReasonText, levelReason, levelReasonText)
		if err != nil {
			return
		}
	}

	if len(newItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
		flag := inventoryManager.BatchAddOfItemLevel(newItemList, itemGetReason, itemGetReasonText)
		if !flag {
			panic("welfare:增加物品应该成功")
		}
	}

	// 更新次数
	info.AttendTimes += attendTimes
	if info.LeftTimes > totalCostTimes {
		info.LeftTimes -= totalCostTimes
	} else {
		info.LeftTimes = 0
	}
	welfareManager.UpdateObj(obj)
	eventData := welfareeventtypes.CreatePlayerAttendDrewEventData(groupId, attendTimes)
	gameevent.Emit(welfareeventtypes.EventTypeAttendDrew, pl, eventData)

	// 添加日志
	for _, itemData := range totalItemList {
		drewLogEventData := welfareeventtypes.CreateDrewAddLogEventData(pl.GetName(), itemData.ItemId, itemData.Num)
		gameevent.Emit(welfareeventtypes.EventTypeDrewAddLog, groupId, drewLogEventData)
	}

	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	logList := welfare.GetWelfareService().GetDrewLogByTime(groupId, lastLogTime)
	scMsg := pbutil.BuildSCOpenActivityCostDrewAttend(dropItemList, extraItemList, logList, groupId, info.Ratio, int32(drewType), dropLevel, rewIndexList)
	pl.SendMsg(scMsg)
	return
}

func countDropItem(pl player.Player, info *drewcostdrewtypes.LuckyCostDrewInfo, groupId, dropLevel int32) (rewIndex int32, dropData *droptemplate.DropItemData) {
	luckTemp := welfaretemplate.GetWelfareTemplateService().GetLuckDrewTemplateByArg(groupId, dropLevel)
	timesMap := luckTemp.GetRewDropByTimesMap()
	timesDescList := luckTemp.GetTimesDesc()

	// 计算物品
	curAttendNum := info.AttendTimes
	curAttendNum += 1
	dropId := luckTemp.DropId
	for _, times := range timesDescList {
		ret := curAttendNum % int32(times)
		if ret == 0 {
			dropId = timesMap[int32(times)]
			break
		}
	}

	flag := droptemplate.GetDropTemplateService().CheckSureDrop(dropId)
	if !flag {
		panic(fmt.Errorf("welfare:掉落包不是必定掉落，dropId:%d", dropId))
	}

	rewIndex, dropData = droptemplate.GetDropTemplateService().GetDropItemLevelWithIndex(dropId)
	if dropData == nil {
		log.WithField("dropId", dropId).Warn("掉落包随机为空")
		return
	}
	dropData.BindType = itemtypes.ItemBindTypeUnBind

	// //道具公告
	// itemId := dropData.GetItemId()
	// num := dropData.GetNum()
	// inventorylogic.PrecioustemBroadcast(pl, itemId, num, lang.InventoryLuckyCostTrayItemNotice)

	if info.Ratio > 0 {
		// 奖励倍数
		dropData.Num = dropData.Num * info.Ratio
		// 重置规则CD
		info.ResetRule()
	}

	return
}
