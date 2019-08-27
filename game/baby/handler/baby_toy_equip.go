package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	babylogic "fgame/fgame/game/baby/logic"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytypes "fgame/fgame/game/baby/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_EQUIP_TOY_TYPE), dispatch.HandlerFunc(handleUseToy))
}

//使用玩具
func handleUseToy(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理使用玩具")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSBabyEquipToy)

	index := csMsg.GetIndex()
	if index < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用宝宝玩具,索引错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	suitType := babytypes.ToySuitType(csMsg.GetSuitType())
	if !suitType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"suitType": suitType,
			}).Warn("inventory:使用宝宝玩具,类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = useToy(tpl, index, suitType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"suitType": suitType,
				"error":    err,
			}).Error("inventory:处理使用玩具,错误")

		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"index":    index,
			"suitType": suitType,
		}).Debug("inventory:处理使用玩具,完成")

	return nil
}

//使用玩具
func useToy(pl player.Player, index int32, suitType babytypes.ToySuitType) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	toyBag := babyManager.GetBabyToyBag(suitType)
	if toyBag == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"suitType": suitType,
			}).Warn("inventory:使用玩具,玩具套装背包不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	it := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
	//物品不存在
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用玩具,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	//判断物品是否可以装备
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	toyTemplate := itemTemplate.GetBabyToyTemplate()
	if toyTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"itemId":   itemId,
				"itemType": itemTemplate.GetItemType(),
			}).Warn("inventory:使用玩具,此物品不是玩具")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotEquip)
		return
	}

	//判断是否已经装备
	pos := toyTemplate.GetPosType()
	if !pos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"pos":      pos,
			}).Warn("inventory:使用玩具,没有可装备的位置")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	toySlot := toyBag.GetByPosition(pos)
	if toySlot != nil && !toySlot.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("inventory:使用玩具,该玩具已经装备")
		playerlogic.SendSystemMessage(pl, lang.BabyToySlotHadEquip)
		return
	}

	//移除物品
	useReason := commonlog.InventoryLogReasonBabyEquipToy
	flag, _ := inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, index, 1, useReason, useReason.String())
	if !flag {
		panic("inventory:使用玩具移除物品应该是可以的")
	}

	flag = toyBag.PutOn(itemId, pos)
	if !flag {
		panic(fmt.Errorf("inventory:穿上位置 [%s]应该是可以的", pos.String()))
	}

	//同步改变
	babylogic.BabyPropertyChanged(pl)
	babylogic.SnapBabyToyChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCBabyEquipToy(index, int32(suitType))
	pl.SendMsg(scMsg)
	return
}
