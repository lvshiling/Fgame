package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/center/center"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TAKE_OUT_ALLIANCE_DEPOT_TYPE), dispatch.HandlerFunc(handleTakeOutDepot))
}

//处理仙盟仓库取出
func handleTakeOutDepot(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟仓库取出")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSTakeOutAllianceDepot)
	index := csMsg.GetIndex()
	num := csMsg.GetNum()
	itemId := csMsg.GetItemId()
	err = takeOutAllianceDepot(tpl, index, itemId, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"error":    err,
			}).Error("alliance:处理仙盟仓库取出,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"index":    index,
		}).Debug("alliance:处理仙盟仓库取出,完成")
	return nil

}

//仙盟仓库取出
func takeOutAllianceDepot(pl player.Player, depotIndex, itemId, num int32) (err error) {
	if !center.GetCenterService().IsAllianceOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:仙盟仓库关闭中")
		playerlogic.SendSystemMessage(pl, lang.AllianceDepotClose)
		return
	}
	allianceId := pl.GetAllianceId()
	al := alliance.GetAllianceService().GetAlliance(allianceId)
	if al == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理仙盟仓库取出,仙盟不存在")
		playerlogic.SendSystemMessage(pl, lang.AllianceNoExist)
		return
	}
	it := al.GetDepotItemByIndex(depotIndex)
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"depotIndex": depotIndex,
			}).Warn("alliance:处理仙盟仓库取出,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	curItemId := it.GetItemId()
	curNum := it.GetNum()
	level := it.GetLevel()
	bind := it.GetBindType()
	propertyData := it.GetPropertyData()

	if curItemId != itemId {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"depotIndex": depotIndex,
				"curNum":     curNum,
				"num":        num,
			}).Warn("alliance:处理仙盟仓库取出,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
	}

	if curNum < num {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"depotIndex": depotIndex,
				"curNum":     curNum,
				"num":        num,
			}).Warn("alliance:处理仙盟仓库取出,超过当前位置最大数量")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"depotIndex": depotIndex,
			}).Warn("alliance:处理仙盟仓库取出,物品数据模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否能取出
	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	costPoint := itemTemp.UnionUse * num
	if !allianceManager.HasEnoughPoint(costPoint) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"depotIndex": depotIndex,
				"costPoint":  costPoint,
			}).Warn("alliance:处理仙盟仓库取出,积分不足")
		playerlogic.SendSystemMessage(pl, lang.AllianceDepotHasNotEnoughPoint)
		return
	}

	// 背包空间
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotItemLevelWithProperty(itemId, num, level, bind, propertyData.GetExpireType(), propertyData.GetItemGetTime(), propertyData.GetExpireTimestamp()) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"depotIndex": depotIndex,
				"itemId":     itemId,
				"num":        num,
			}).Warn("alliance:处理仙盟仓库取出,背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	// 仓库移除
	err = alliance.GetAllianceService().TakeOutDepot(pl.GetAllianceId(), depotIndex, itemId, num)
	if err != nil {
		return
	}

	// 背包添加
	reason := commonlog.InventoryLogReasonTakeOutAllianceDepot
	flag := inventoryManager.AddItemLevelWithPropertyData(itemId, num, level, bind, propertyData, reason, reason.String())
	if !flag {
		panic(fmt.Errorf("alliance:仙盟仓库取出物品，添加物品应该成功,itemId:%d", itemId))
	}

	// 消耗积分
	beforePoint := allianceManager.GetDepotPoint()
	allianceManager.CostDepotPoint(costPoint)
	inventorylogic.SnapInventoryChanged(pl)

	// 仓库取出日志
	removeItemReason := commonlog.AllianceLogReasonDepotTakeOutItem
	removeItemReasonText := fmt.Sprintf(removeItemReason.String(), pl.GetId())
	depotLogEventData := allianceeventtypes.CreateAllianceDepotItemChangedLogEventData(itemId, num, removeItemReason, removeItemReasonText)
	gameevent.Emit(allianceeventtypes.EventTypeAllianceDepotItemChangedLog, al, depotLogEventData)

	//玩家积分变化日志
	pointReason := commonlog.AllianceLogReasonPlayerDepotPointChanged
	pointReasonText := fmt.Sprintf(pointReason.String(), itemId)
	pointLogEventData := allianceeventtypes.CreatePlayerAllianceDepotPointLogEventData(beforePoint, costPoint, pointReason, pointReasonText)
	gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceDepotPointChangedLog, al, pointLogEventData)

	//广播帮派
	format := lang.GetLangService().ReadLang(lang.AllianceDepotTakeOutNotice)
	itemName := coreutils.FormatColor(itemTemp.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemp.FormateItemNameOfNum(num)))
	linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
	itemNameLink := coreutils.FormatLink(itemName, linkArgs)
	takerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, pl.GetName())

	args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(funcopentypes.FuncOpenTypeAlliance), al.GetAllianceId()}
	link := coreutils.FormatLink(chattypes.ButtonTypeToLook, args)
	content := fmt.Sprintf(format, takerName, itemNameLink, link)
	chatlogic.BroadcastAllianceSystem(al.GetAllianceId(), pl.GetId(), pl.GetName(), chattypes.MsgTypeText, []byte(content), "")

	curPoint := allianceManager.GetDepotPoint()
	scMsg := pbutil.BuildSCTakeOutAllianceDepot(curPoint)
	pl.SendMsg(scMsg)
	return
}
