package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/core/utils"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_ITEM_DECOMPOSE_TYPE), dispatch.HandlerFunc(handleInventoryItemDecompose))
}

// 物品拆解
func handleInventoryItemDecompose(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理背包拆解消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSInventoryItemDecompose)
	index := csMsg.GetIndex()
	num := csMsg.GetNum()

	if num <= 0 || index < 0 {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"index":    index,
				"num":      num,
			}).Warn("inventory:处理背包拆解,参数不对")
		return
	}
	bagType := inventorytypes.BagTypePrim
	bagTypePtr := csMsg.BagType
	if bagTypePtr != nil {
		bagType = inventorytypes.BagType(csMsg.GetBagType())
		if !bagType.Valid() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"bagType":  bagType,
				}).Warn("inventory:处理背包拆解,参数不对")
			playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
			return
		}
	}

	err = decomposeItem(tpl, bagType, index, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"index":    index,
				"num":      num,
			}).Error("inventory:处理背包拆解,错误")

		return
	}
	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"index":    index,
			"num":      num,
		}).Debug("inventory:处理背包拆解消息,成功")

	return
}

func decomposeItem(pl player.Player, bagType inventorytypes.BagType, index int32, num int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	it := inventoryManager.FindItemByIndex(bagType, index)
	if it == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"bagType":  bagType,
			}).Warn("inventory:拆解物品，物品不存")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"bagType":  bagType,
				"itemId":   itemId,
			}).Warn("inventory:拆解物品，物品不存")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//拆解数量校验
	if num > it.Num {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"bagType":  bagType,
				"curNum":   it.Num,
				"needNum":  num,
			}).Warn("inventory:拆解物品，物品数量不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//能否拆解
	chaiJieItemMap := itemTemplate.GetChaiJieItemMap()
	if !itemTemplate.IsCanChaiJie() {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"index":          index,
				"bagType":        bagType,
				"itemId":         itemId,
				"chaiJieItemMap": chaiJieItemMap,
			}).Warn("inventory:拆解物品，物品不能拆解")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotChaiJie)
		return
	}

	chaiJieItemMap = utils.MultMap(chaiJieItemMap, num)
	if !inventoryManager.HasEnoughSlots(chaiJieItemMap) {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"index":          index,
				"bagType":        bagType,
				"itemId":         itemId,
				"num":            num,
				"chaiJieItemMap": chaiJieItemMap,
			}).Warn("inventory:拆解物品，背包不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	useReason := commonlog.InventoryLogReasonItemChaiJieCost
	useReasonText := fmt.Sprintf(useReason.String(), itemId)
	flag, _ := inventoryManager.RemoveIndex(bagType, index, num, useReason, useReasonText)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"itemId":   itemId,
				"num":      num,
			}).Warn("inventory:拆解物品失败")
		return
	}

	getReason := commonlog.InventoryLogReasonItemChaiJieGet
	getReasonText := fmt.Sprintf(useReason.String(), itemId)
	if flag := inventoryManager.BatchAdd(chaiJieItemMap, getReason, getReasonText); !flag {
		panic(fmt.Errorf("inventory：拆解获得物品应该成功"))
	}

	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCInventoryItemDecompose(bagType, index, num, chaiJieItemMap)
	pl.SendMsg(scMsg)
	return
}
