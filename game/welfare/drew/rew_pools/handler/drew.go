package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	commonlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
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
	rewpoolstemplate "fgame/fgame/game/welfare/drew/rew_pools/template"
	rewpoolstypes "fgame/fgame/game/welfare/drew/rew_pools/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_REW_POOLS_DREW_TYPE), dispatch.HandlerFunc(rewPoolsDrewHandler))
}

func rewPoolsDrewHandler(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityRewPoolsDrew)
	groupId := csMsg.GetGroupId()
	isAutoBuy := csMsg.GetIsAutoBuy()
	attendType := drewcommontypes.LuckyDrewAttendType(csMsg.GetAttendType())

	if !attendType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"attendType": attendType,
			}).Warn("welfare:处理奖池抽奖,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = rewPoolsDrew(tpl, groupId, isAutoBuy, attendType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("welfare:处理奖池抽奖，错误")
		return
	}
	return
}

func rewPoolsDrew(pl player.Player, groupId int32, isAutoBuy bool, attendType drewcommontypes.LuckyDrewAttendType) (err error) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeRewPools

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:奖池抽奖错误，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:奖池抽奖错误，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	groupTemp := groupInterface.(*rewpoolstemplate.GroupTemplateRewPools)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	info := obj.GetActivityData().(*rewpoolstypes.RewPoolsInfo)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	welfaretemplateService := welfaretemplate.GetWelfareTemplateService()
	luckDrewTemp := welfaretemplateService.GetLuckDrewTemplate(groupId)
	if luckDrewTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:奖池抽奖错误，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	loop := attendType.GetAttendNum()
	if inventoryManager.GetEmptySlots(inventorytypes.BagTypePrim) < loop {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:奖池抽奖错误，空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnoughSlot, fmt.Sprintf("%d", loop))
		return
	}
	//检查
	needGold := int64(0)
	needBindGold := int64(0)
	needSilver := int64(0)
	outsideFinalShopIdMap := make(map[int32]int32)
	useItemMap := make(map[int32]int32)
	// 都是使用同样抽奖卷数量
	for itemId, itemNeedNum := range luckDrewTemp.GetUseItemMap() {
		itemNeedNum = itemNeedNum * attendType.GetAttendNum()
		curItemNum := inventoryManager.NumOfItems(itemId)
		finalUseItemNum := itemNeedNum
		isEnoughBuyTimes := true
		shopIdMap := make(map[int32]int32)
		if curItemNum < itemNeedNum {
			if !isAutoBuy {
				log.WithFields(
					log.Fields{
						"playerId":    pl.GetId(),
						"groupId":     groupId,
						"itemId":      itemId,
						"itemNeedNum": itemNeedNum,
					}).Warn("welfare:奖池抽奖错误，数量不足")
				playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
				return
			}
			needBuyNum := itemNeedNum - curItemNum
			finalUseItemNum = curItemNum
			if needBuyNum > 0 {
				if !shop.GetShopService().ShopIsSellItem(itemId) {
					log.WithFields(log.Fields{
						"playerId":  pl.GetId(),
						"useItemId": itemId,
						"autoFlag":  isAutoBuy,
					}).Warn("welfare:商铺没有该道具,无法自动购买")
					playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
					return
				}

				isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, itemId, needBuyNum)
				if !isEnoughBuyTimes {
					log.WithFields(log.Fields{
						"playerId":  pl.GetId(),
						"useItemId": itemId,
						"autoFlag":  isAutoBuy,
					}).Warn("welfare:购买物品失败，奖池抽奖")
					playerlogic.SendSystemMessage(pl, lang.ShopBuyNumInvalid)
					return
				}

				shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
				needGold += shopNeedGold
				needBindGold += shopNeedBindGold
				needSilver += shopNeedSilver
			}
		}

		outsideFinalShopIdMap = coreutils.MergeMap(outsideFinalShopIdMap, shopIdMap)
		if finalUseItemNum > 0 {
			useItemMap[itemId] = finalUseItemNum
		}
	}

	//是否足够银两
	flag := propertyManager.HasEnoughSilver(needSilver)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:奖池抽奖，银两不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//是否足够元宝
	flag = propertyManager.HasEnoughGold(needGold, false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:奖池抽奖，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//是否足够绑元
	needCostBindGold := needBindGold + needGold
	flag = propertyManager.HasEnoughGold(needCostBindGold, true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:奖池抽奖，绑元不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//奖池进阶或回退处理
	dropIdList := []int32{}
	loopTimes := groupTemp.GetFirstValue1()
	if loopTimes <= 0 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"loopTimes": loopTimes,
			}).Warn("welfare:奖池抽奖错误，回退最大次数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	pathRecord := []commontypes.PathType{}
	backTimesRecord := info.GetBackTimes()
	postionRecord := info.Position

	rewlist := []*commontypes.RewNode{}
	for _, luckTemp := range groupTemp.GetLuckDrewTemp() {
		rewNode := commontypes.CreateRewNode(luckTemp.Level, luckTemp.Percent1, luckTemp.Percent2, luckTemp.Rate)
		rewlist = append(rewlist, rewNode)
	}
	rewpools := commontypes.CreateRewPools(rewlist)
	_, _, pathRecord, err = commonlogic.SimulateRewPools(rewpools, postionRecord, loop, backTimesRecord, loopTimes)
	if err != nil {
		return err
	}

	// 奖池奖励
	postionRecord = info.Position
	for i := int32(0); i < int32(len(pathRecord)); i++ {
		pathTyp := pathRecord[i]
		switch pathTyp {
		case commontypes.PathTypeBackTimesEnoughForward:
			luckDrewTemp = welfaretemplateService.GetLuckDrewTemplateByArg(groupId, postionRecord)
			dropId := luckDrewTemp.DropId
			dropIdList = append(dropIdList, dropId)
			postionRecord++
		case commontypes.PathTypeBack:
			luckDrewTemp = welfaretemplateService.GetLuckDrewTemplateByArg(groupId, postionRecord)
			dropId := luckDrewTemp.DropId
			dropIdList = append(dropIdList, dropId)
			postionRecord--
		case commontypes.PathTypeForward:
			luckDrewTemp = welfaretemplateService.GetLuckDrewTemplateByArg(groupId, postionRecord)
			dropId := luckDrewTemp.DropId
			dropIdList = append(dropIdList, dropId)
			postionRecord++
		case commontypes.PathTypeStill:
			luckDrewTemp = welfaretemplateService.GetLuckDrewTemplateByArg(groupId, postionRecord)
			dropId := luckDrewTemp.DropId
			dropIdList = append(dropIdList, dropId)
		default:
			continue
		}
	}

	//获取奖励信息
	totalItemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)
	for _, data := range totalItemList {
		if data == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"groupId":  groupId,
				}).Warn("welfare:奖池抽奖错误，掉落包不存在")
			playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(outsideFinalShopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, outsideFinalShopIdMap)
	}

	newTotalItemList, resMap := droplogic.SeperateItemDatas(totalItemList)

	//背包空间
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newTotalItemList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:领取活动奖励请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//----------------------------------------检查&模拟 和 消耗&实际行动 分割线--------------------------------------

	//自动购买消耗金钱
	//消耗银两
	if needSilver > 0 {
		silverUseReason := commonlog.SilverLogReasonRewPools
		silverUseReasonText := fmt.Sprintf(silverUseReason.String(), info.Position)
		flag := propertyManager.CostSilver(needSilver, silverUseReason, silverUseReasonText)
		if !flag {
			panic("welfare:消耗银两应该成功")
		}
	}

	//消耗元宝
	if needGold > 0 {
		goldUseReason := commonlog.GoldLogReasonDrewPools
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), info.Position)
		flag := propertyManager.CostGold(needGold, false, goldUseReason, goldUseReasonText)
		if !flag {
			panic("welfare:消耗元宝应该成功")
		}
	}

	//消耗绑元
	if needBindGold > 0 {
		goldUseReason := commonlog.GoldLogReasonDrewPools
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), info.Position)
		flag := propertyManager.CostGold(needBindGold, true, goldUseReason, goldUseReasonText)
		if !flag {
			panic("welfare:消耗元宝应该成功")
		}
	}

	//消耗抽奖卷
	if len(useItemMap) > 0 {
		useReason := commonlog.InventoryLogReasonRewPoolsDrew
		useReasonText := fmt.Sprintf(useReason.String(), info.Position)
		flag = inventoryManager.BatchRemove(useItemMap, useReason, useReasonText)
		if !flag {
			panic("inventory:移除物品应该是可以的")
		}
	}

	//获取奖励
	if len(newTotalItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonRewPoolsDrew
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), info.Position)
		flag := inventoryManager.BatchAddOfItemLevel(newTotalItemList, itemGetReason, itemGetReasonText)
		if !flag {
			panic("inventory:增加物品应该成功")
		}
	}
	goldGetReason := commonlog.GoldLogReasonRewPoolsDrew
	goldGetReasonText := fmt.Sprintf(goldGetReason.String())
	silverGetReason := commonlog.SilverLogReasonRewPoolsDrew
	silverGetReasonText := fmt.Sprintf(silverGetReason.String())
	levelGetReason := commonlog.LevelLogReasonRewPoolsDrew
	levelGetReasonText := fmt.Sprintf(levelGetReason.String())
	if len(resMap) > 0 {
		droplogic.AddRes(pl, resMap, goldGetReason, goldGetReasonText, silverGetReason, silverGetReasonText, levelGetReason, levelGetReasonText)
	}
	//奖池位置改变
	for _, pathTyp := range pathRecord {
		switch pathTyp {
		case commontypes.PathTypeBackTimesEnoughForward:
			info.BackTimesEnoughPoolForWard(loopTimes)
		case commontypes.PathTypeBack:
			info.PoolBack()
		case commontypes.PathTypeForward:
			info.PoolForWard()
		case commontypes.PathTypeStill:
			continue
		default:
			continue
		}
	}

	// 添加日志
	for _, itemData := range totalItemList {
		drewLogEventData := welfareeventtypes.CreateDrewAddLogEventData(pl.GetName(), itemData.ItemId, itemData.Num)
		gameevent.Emit(welfareeventtypes.EventTypeDrewAddLog, groupId, drewLogEventData)
	}
	logList := welfare.GetWelfareService().GetDrewLogByTime(groupId, 0)

	welfareManager.UpdateObj(obj)
	//更新次数
	attendNum := attendType.GetAttendNum()
	eventData := welfareeventtypes.CreatePlayerAttendDrewEventData(groupId, attendNum)
	gameevent.Emit(welfareeventtypes.EventTypeAttendDrew, pl, eventData)

	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityRewPoolsDrew(groupId, info.Position, totalItemList, info.BackTimes, logList)
	pl.SendMsg(scMsg)
	return
}
