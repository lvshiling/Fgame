package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MIBAO_DEPOT_TAKE_OUT_TYPE), dispatch.HandlerFunc(handleMiBaoDepotTakeOut))
}

//处理秘宝仓库取出物品
func handleMiBaoDepotTakeOut(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理秘宝仓库取出物品")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMiBaoDepotTakeOut := msg.(*uipb.CSMibaoDepotTakeOut)
	typ := equipbaokutypes.BaoKuType(csMiBaoDepotTakeOut.GetType())
	isBatch := csMiBaoDepotTakeOut.GetIsBatch()
	index := int32(0)
	if !isBatch {
		index = csMiBaoDepotTakeOut.GetIndex()
	}

	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"type":     typ,
				"error":    err,
			}).Error("inventory:处理秘宝仓库取出物品,仓库类型不合法")
		return
	}

	err = miBaoDepotTakeOut(tpl, typ, isBatch, index)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"type":     typ,
				"error":    err,
			}).Error("inventory:处理秘宝仓库取出物品,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("inventory:处理秘宝仓库取出物品,完成")
	return nil
}

//秘宝仓库取出物品
func miBaoDepotTakeOut(pl player.Player, typ equipbaokutypes.BaoKuType, isBatch bool, index int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	leftSlotNum := inventoryManager.GetEmptySlots(inventorytypes.BagTypePrim)
	if leftSlotNum == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("inventory:秘宝仓库取出失败, 背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	indexList := inventoryManager.FindMiBaoDepotItemIndexsFromEnd(typ)
	if indexList == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("inventory:秘宝仓库取出失败，没有物品取出")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	// 是否批量取出
	if isBatch {
		for _, itemIndex := range indexList {
			itemObj := inventoryManager.FindMiBaoDepotItemByIndex(itemIndex, typ)
			if itemObj == nil || itemObj.IsEmpty() {
				continue
			}

			itemId := itemObj.ItemId
			itemNum := itemObj.Num
			level := itemObj.Level
			bind := itemObj.BindType
			propertyData := itemObj.PropertyData
			//背包空间是否足够
			flag := inventoryManager.HasEnoughSlotItemLevelWithProperty(itemId, itemNum, level, bind, propertyData.GetExpireType(), propertyData.GetItemGetTime(), propertyData.GetExpireTimestamp())
			if !flag {
				break
			}

			flag, _ = inventoryManager.RemoveMiBaoDepotByIndex(itemIndex, itemNum, typ)
			if !flag {
				panic("inventory: 秘宝仓库移除物品应该成功")
			}

			itemUseReason := commonlog.InventoryLogReasonTakeOutMiBaoDepot
			flag = inventoryManager.AddItemLevelWithPropertyData(itemId, itemNum, level, bind, propertyData, itemUseReason, itemUseReason.String())
			if !flag {
				panic("inventory:存入背包应该成功")
			}
		}
	} else {
		itemObj := inventoryManager.FindMiBaoDepotItemByIndex(index, typ)
		if itemObj == nil || itemObj.IsEmpty() {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"itemIndex": index,
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
					"itemIndex": index,
					"itemId":    itemId,
					"itemNum":   itemNum,
				}).Warn("inventory:仓库取出失败，背包空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}

		flag, _ = inventoryManager.RemoveMiBaoDepotByIndex(index, itemNum, typ)
		if !flag {
			panic("inventory: 秘宝仓库移除物品应该成功")
		}

		itemUseReason := commonlog.InventoryLogReasonTakeOutMiBaoDepot
		flag = inventoryManager.AddItemLevelWithPropertyData(itemId, itemNum, level, bind, propertyData, itemUseReason, itemUseReason.String())
		if !flag {
			panic("inventory:存入背包应该成功")
		}
	}

	inventorylogic.SnapInventoryChanged(pl)
	inventorylogic.SnapMiBaoDepotChanged(pl, typ)

	itemChangedList := inventoryManager.GetMiBaoDepotChangedSlotAndReset(typ)
	scDepotTakeOut := pbutil.BuildSCMiBaoDepotTakeOut(itemChangedList, int32(typ), isBatch, index)
	pl.SendMsg(scDepotTakeOut)
	return
}
