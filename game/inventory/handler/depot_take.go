package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DEPOT_TAKE_OUT_TYPE), dispatch.HandlerFunc(handleDepotTakeOut))
}

//处理仓库取出物品
func handleDepotTakeOut(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理仓库取出物品")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csDepotTakeOut := msg.(*uipb.CSDepotTakeOut)
	itemIndex := csDepotTakeOut.GetIndex()

	err = depotTakeOut(tpl, itemIndex)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"itemIndex": itemIndex,
				"error":     err,
			}).Error("inventory:处理仓库取出物品,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"itemIndex": itemIndex,
		}).Debug("inventory:处理仓库取出物品,完成")
	return nil
}

//仓库取出物品
func depotTakeOut(pl player.Player, itemIndex int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemObj := inventoryManager.FindDepotItemByIndex(itemIndex)
	if itemObj == nil || itemObj.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"itemIndex": itemIndex,
			}).Warn("inventory:仓库取出失败，物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	itemId := itemObj.ItemId
	itemNum := itemObj.Num
	level := itemObj.Level
	bind := itemObj.BindType
	propertyData := itemObj.PropertyData

	//背包空间是否足够
	flag := inventoryManager.HasEnoughSlotItemLevelWithProperty(itemId, itemNum, level, bind, propertyData.GetExpireType(), propertyData.GetItemGetTime(), propertyData.GetExpireTimestamp())
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"itemIndex": itemIndex,
				"itemId":    itemId,
				"itemNum":   itemNum,
			}).Warn("inventory:仓库取出失败，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	flag, _ = inventoryManager.RemoveDepotByIndex(itemIndex, itemNum)
	if !flag {
		panic("inventory: 仓库移除物品应该成功")
	}

	itemUseReason := commonlog.InventoryLogReasonTakeOutDepot
	flag = inventoryManager.AddItemLevelWithPropertyData(itemId, itemNum, level, bind, propertyData, itemUseReason, itemUseReason.String())
	if !flag {
		panic("inventory:存入背包应该成功")
	}

	inventorylogic.SnapInventoryChanged(pl)
	inventorylogic.SnapDepotChanged(pl)

	itemChangedList := inventoryManager.GetDepotChangedSlotAndReset()
	scDepotTakeOut := pbutil.BuildSCDepotTakeOut(itemChangedList)
	pl.SendMsg(scDepotTakeOut)
	return
}
