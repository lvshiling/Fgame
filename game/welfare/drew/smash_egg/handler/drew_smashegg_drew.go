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
	drewsmasheggtemplate "fgame/fgame/game/welfare/drew/smash_egg/template"
	drewsmasheggtypes "fgame/fgame/game/welfare/drew/smash_egg/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_SMASH_EGG_ATTEND_TYPE), dispatch.HandlerFunc(handleAttendSmashEgg))

}

//砸金蛋
func handleAttendSmashEgg(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:砸金蛋")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csSmashEggAttend := msg.(*uipb.CSOpenActivitySmashEggAttend)
	groupId := csSmashEggAttend.GetGroupId()
	attendCount := csSmashEggAttend.GetAttendCount()
	dropNum := csSmashEggAttend.GetDropNum()

	err = attendSmashEgg(tpl, groupId, attendCount, dropNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("welfare:处理砸金蛋,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理砸金蛋完成")
	return nil

}

//砸金蛋逻辑
func attendSmashEgg(pl player.Player, groupId, attendNum, dropNum int32) (err error) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeSmashEgg

	if dropNum < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"dropNum":  dropNum,
			}).Warn("welfare:砸金蛋抽奖请求，抽奖号码错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:砸金蛋请求错误，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:砸金蛋请求错误，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	groupTemp := groupInterface.(*drewsmasheggtemplate.GroupTemplateSmashEgg)
	batchMaxCount := groupTemp.GetSmashEggBatchCount()

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*drewsmasheggtypes.SmashEggInfo)

	if batchMaxCount <= int32(len(info.AttendTimesList)) {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"batchMaxCount": batchMaxCount,
				"batchCurCount": int32(len(info.AttendTimesList)),
			}).Warn("welfare:砸金蛋抽奖请求，数据信息错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if dropNum == 0 {
		attendNum = batchMaxCount - int32(len(info.AttendTimesList))
	} else {
		attendNum = 1
	}

	if attendNum > 0 {
		if inventoryManager.GetEmptySlots(inventorytypes.BagTypePrim) < attendNum {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("welfare:充值抽奖错误，空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnoughSlot, fmt.Sprintf("%d", attendNum))
			return
		}
	}

	//是否足够钱
	needGold := groupTemp.GetNeedGold() * attendNum
	needBindGold := groupTemp.GetNeedBindGold() * attendNum
	needSilver := groupTemp.GetNeedSilver() * attendNum
	flag := propertyManager.HasEnoughCost(int64(needBindGold), int64(needGold), int64(needSilver))
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("welfare:砸金蛋错误，绑元等资源不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	luckTemp := welfaretemplate.GetWelfareTemplateService().GetLuckDrewTemplate(groupId)
	if luckTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:砸金蛋错误，抽奖模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	timesMap := luckTemp.GetRewDropByTimesMap()
	timesDescList := luckTemp.GetTimesDesc()

	// 计算物品
	var totalItemList []*droptemplate.DropItemData
	var dropItemList []*droptemplate.DropItemData
	var extraItemList []*droptemplate.DropItemData
	curAttendNum := info.AttendTimes
	for index := int32(0); index < attendNum; index++ {
		curAttendNum += 1
		dropId := luckTemp.DropId
		for _, times := range timesDescList {
			ret := curAttendNum % int32(times)
			if ret == 0 {
				dropId = timesMap[int32(times)]
				break
			}
		}

		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
		if dropData == nil {
			log.WithField("dropId", dropId).Warn("掉落包随机为空")
			continue
		}
		dropData.BindType = itemtypes.ItemBindTypeUnBind
		totalItemList = append(totalItemList, dropData)
		dropItemList = append(dropItemList, dropData)
		//额外奖励
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
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"len":      len(newItemList),
			}).Warn("welfare:砸金蛋错误，空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗钱
	goldUseReason := commonlog.GoldLogReasonDrewUse
	silverUseReason := commonlog.SilverLogReasonDrewUse
	goldUseReasonText := fmt.Sprintf(goldUseReason.String(), subType)
	silverUseReasonText := fmt.Sprintf(silverUseReason.String(), subType)
	flag = propertyManager.Cost(int64(needBindGold), int64(needGold), goldUseReason, goldUseReasonText, int64(needSilver), silverUseReason, silverUseReasonText)
	if !flag {
		panic("welfare:消耗元宝应该成功")
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
		flag = inventoryManager.BatchAddOfItemLevel(newItemList, itemGetReason, itemGetReasonText)
		if !flag {
			panic("welfare:增加物品应该成功")
		}
	}

	// 更新次数
	info.AttendTimes += attendNum
	if dropNum > 0 && batchMaxCount > int32(len(info.AttendTimesList)+1) {
		info.AttendTimesList = append(info.AttendTimesList, dropNum)
	} else {
		info.AttendTimesList = nil
	}
	welfareManager.UpdateObj(obj)
	// eventData := welfareeventtypes.CreatePlayerAttendDrewEventData(groupId, attendNum)
	// gameevent.Emit(welfareeventtypes.EventTypeAttendDrew, pl, eventData)

	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scSmashEggAttend := pbutil.BuildSCOpenActivitySmashEggAttend(dropItemList, extraItemList, groupId, attendNum, dropNum)
	pl.SendMsg(scSmashEggAttend)
	return
}
