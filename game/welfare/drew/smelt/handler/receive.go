package handler

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	smelttypes "fgame/fgame/game/welfare/drew/smelt/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	pbutil "fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveMultipleHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmelt, welfare.ReceiveMultipleHandlerFunc(smeltReceiveHandler))
}

func smeltReceiveHandler(pl player.Player, rewId int32, receiveType welfaretypes.ReceiveType) (err error) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeSmelt
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取冶炼奖励请求，模板不存在")
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
			}).Warn("welfare:领取冶炼奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*smelttypes.SmeltInfo)
	needItemNum := openTemp.Value2
	if needItemNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"needItemNum": needItemNum,
			}).Warn("welfare:领取冶炼奖励请求，配置所需物品的值Value2不正确")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	remainTimes := info.GetRemainCanReceiveRecord(needItemNum)
	needTimes := receiveType.ToInt32()
	if remainTimes < needTimes {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"needItemNum": needItemNum,
				"Num":         info.Num,
			}).Warn("welfare:领取冶炼奖励请求，消耗数量不足或者已经领取")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//运营抽奖模板
	luckDrewTemp := welfaretemplate.GetWelfareTemplateService().GetLuckDrewTemplate(groupId)
	if luckDrewTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:奖池抽奖错误，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if inventoryManager.GetEmptySlots(inventorytypes.BagTypePrim) < receiveType.ToInt32() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:奖池抽奖错误，空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnoughSlot, fmt.Sprintf("%d", receiveType.ToInt32()))
		return
	}

	dropId := luckDrewTemp.DropId
	dropIdList := []int32{}
	for i := int32(0); i < receiveType.ToInt32(); i++ {
		dropIdList = append(dropIdList, dropId)
	}

	//获取奖励信息
	totalItemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)

	if len(dropIdList) != len(totalItemList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:奖池抽奖错误，掉落包缺少")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	newTotalItemList, resMap := droplogic.SeperateItemDatas(totalItemList)

	//获取奖励
	if len(newTotalItemList) > 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//背包空间
		if !inventoryManager.HasEnoughSlotsOfItemLevel(newTotalItemList) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("welfare:领取活动奖励请求，背包空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
		flag := inventoryManager.BatchAddOfItemLevel(newTotalItemList, itemGetReason, itemGetReasonText)
		if !flag {
			panic("inventory:增加物品应该成功")
		}
	}

	goldGetReason := commonlog.GoldLogReasonOpenActivityRew
	goldGetReasonText := fmt.Sprintf(goldGetReason.String(), typ, subType)
	silverGetReason := commonlog.SilverLogReasonOpenActivityRew
	silverGetReasonText := fmt.Sprintf(silverGetReason.String(), typ, subType)
	levelGetReason := commonlog.LevelLogReasonOpenActivityRew
	levelGetReasonText := fmt.Sprintf(levelGetReason.String(), typ, subType)
	if len(resMap) > 0 {
		droplogic.AddRes(pl, resMap, goldGetReason, goldGetReasonText, silverGetReason, silverGetReasonText, levelGetReason, levelGetReasonText)
	}
	rewData := &propertytypes.RewData{}

	//更新信息
	info.AddReceiveRecord(receiveType.ToInt32())
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)
	record := []int32{info.HasReceiveNum}
	scMsg := pbutil.BuildSCOpenActivityReceiveRewMultiple(rewId, groupId, rewData, totalItemList, record)
	pl.SendMsg(scMsg)
	return
}
