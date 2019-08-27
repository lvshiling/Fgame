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
	gamesession "fgame/fgame/game/session"
	drewcommontypes "fgame/fgame/game/welfare/drew/common/types"
	drewcrazyboxtemplate "fgame/fgame/game/welfare/drew/crazy_box/template"
	drewcrazyboxtypes "fgame/fgame/game/welfare/drew/crazy_box/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_CRAZY_BOX_ATTEND_TYPE), dispatch.HandlerFunc(handleAttendCrazyBox))

}

const (
	defaultNum = int32(1)
	batchNum   = int32(10)
)

//疯狂宝箱抽奖
func handleAttendCrazyBox(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:疯狂宝箱抽奖")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityCrazyBoxAttend)
	groupId := csMsg.GetGroupId()
	drewInt := csMsg.GetDrewType()
	lastLogTime := csMsg.GetLastLogTime()

	drewType := drewcommontypes.LuckyDrewAttendType(drewInt)
	if !drewType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:疯狂宝箱抽奖错误，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = attendCrazyBox(tpl, groupId, drewType, lastLogTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("welfare:处理疯狂宝箱抽奖,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理疯狂宝箱抽奖完成")
	return nil

}

//疯狂宝箱抽奖逻辑
func attendCrazyBox(pl player.Player, groupId int32, drewType drewcommontypes.LuckyDrewAttendType, lastLogTime int64) (err error) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeCrazyBox

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:疯狂宝箱抽奖错误，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:疯狂宝箱抽奖错误，活动模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	groupTemp := groupInterface.(*drewcrazyboxtemplate.GroupTemplateCrazyBox)

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*drewcrazyboxtypes.CrazyBoxInfo)

	curBoxLevel, boxLeftTimes := groupTemp.GetCrazyBoxArg(info.AttendTimes)
	totalTimes := groupTemp.GetCrazyBoxTotalTimes(info.GoldNum)
	leftTimes := totalTimes - info.AttendTimes
	luckTemp := welfaretemplate.GetWelfareTemplateService().GetLuckDrewTemplateByArg(groupId, curBoxLevel)
	if luckTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:疯狂宝箱抽奖错误，抽奖模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	openActivityTemp := groupTemp.GetOpenActivityCrazyBox(curBoxLevel)
	if openActivityTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:疯狂宝箱抽奖错误，等级活动模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if boxLeftTimes <= 0 || leftTimes <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:疯狂宝箱抽奖错误，当前宝箱次数不足")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	attendNum := defaultNum
	if drewType == drewcommontypes.LuckyDrewTypeBatch {
		attendNum = batchNum
	}

	if attendNum > boxLeftTimes {
		attendNum = boxLeftTimes
	}

	if attendNum > leftTimes {
		attendNum = leftTimes
	}

	timesMap := luckTemp.GetRewDropByTimesMap()
	timesDescList := luckTemp.GetTimesDesc()

	// 计算物品
	var totalItemList []*droptemplate.DropItemData
	var dropItemList []*droptemplate.DropItemData
	var extraItemList []*droptemplate.DropItemData
	curAttendNum := groupTemp.GetOpenActivityCrazyBoxUpTimes(curBoxLevel) - boxLeftTimes
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
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
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

	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"len":      len(newItemList),
			}).Warn("welfare:疯狂宝箱奖错误，空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
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
	info.AttendTimes += attendNum
	welfareManager.UpdateObj(obj)
	// //抽奖次数排行榜事件
	// eventData := welfareeventtypes.CreatePlayerAttendDrewEventData(groupId, attendNum)
	// gameevent.Emit(welfareeventtypes.EventTypeAttendDrew, pl, eventData)
	// 添加日志
	for _, itemData := range totalItemList {
		drewLogEventData := welfareeventtypes.CreateCrazyBoxAddLogEventData(pl.GetName(), itemData.ItemId, itemData.Num)
		gameevent.Emit(welfareeventtypes.EventTypeCrazyBoxAddLog, groupId, drewLogEventData)
	}
	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)
	//
	logList := welfare.GetWelfareService().GetCrazyBoxLogByTime(groupId, lastLogTime)
	scMsg := pbutil.BuildSCOpenActivityCrazyBoxAttend(dropItemList, extraItemList, logList, groupId, int32(drewType))
	pl.SendMsg(scMsg)
	return
}
