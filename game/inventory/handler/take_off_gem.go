package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/inventory/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_TAKE_OFF_GEM_TYPE), dispatch.HandlerFunc(handleInventoryTakeOffGem))
}

//处理脱下
func handleInventoryTakeOffGem(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理脱下装备宝石")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryTakeOffGem := msg.(*uipb.CSInventoryTakeOffGem)
	slotId := csInventoryTakeOffGem.GetSlotId()
	slotPosition := inventorytypes.BodyPositionType(slotId)
	//参数不对
	if !slotPosition.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"slotId":   slotId,
			}).Warn("inventory:处理脱下装备宝石,参数不对")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	order := csInventoryTakeOffGem.GetOrder()
	if order < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"order":    order,
			}).Warn("inventory:处理脱下装备宝石,参数不对")
	}

	err = takeOffGem(tpl, slotPosition, order)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理脱下装备宝石,错误")
		return err
	}

	log.Debug("inventory:处理脱下装备宝石,完成")
	return nil
}

//脱下
func takeOffGem(pl player.Player, pos inventorytypes.BodyPositionType, order int32) (err error) {
	flag := takeOffGemInternal(pl, pos, order)
	if !flag {
		return
	}

	//同步改变
	inventorylogic.UpdateEquipmentProperty(pl)
	logic.SnapInventoryEquipChanged(pl)
	logic.SnapInventoryChanged(pl)

	//脱下成功
	scInventoryTakeOffEquip := pbutil.BuildSCInventoryTakeOffEquip(pos)
	pl.SendMsg(scInventoryTakeOffEquip)

	return nil
}

func takeOffGemInternal(pl player.Player, pos inventorytypes.BodyPositionType, order int32) (flag bool) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipmentBag := manager.GetEquipmentBag()
	item := equipmentBag.GetByPosition(pos)
	//没有东西
	if item == nil || item.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("inventory:处理脱下装备宝石,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipGemCanNotTakeOff)
		return
	}
	if !equipmentBag.IfEmbedGem(pos, order) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("inventory:处理脱下装备宝石,已经被卸下")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipGemAlreadyTakeOff)
		return
	}

	//没有足够的空间放脱下来的宝石
	num := int32(1)
	if !manager.HasEnoughSlot(item.ItemId, num) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("inventory:处理脱下装备宝石,背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	itemId := equipmentBag.TakeOffGem(pos, order)
	if itemId == 0 {
		panic(fmt.Errorf("inventory:take off should more than 0"))
	}

	//添加物品
	reasonText := commonlog.InventoryLogReasonTakeOffGem.String()
	flag = manager.AddItem(itemId, num, commonlog.InventoryLogReasonTakeOffGem, reasonText)
	if !flag {
		panic(fmt.Errorf("inventory:add item should be success"))
	}
	return
}
