package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_TAKE_OFF_GEM_TYPE), dispatch.HandlerFunc(handleGoldEquipTakeOffGem))
}

//处理脱下
func handleGoldEquipTakeOffGem(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理脱下装备宝石")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csGoldEquipTakeOffGem := msg.(*uipb.CSGoldEquipTakeOffGem)
	slotId := csGoldEquipTakeOffGem.GetSlotId()
	slotPosition := inventorytypes.BodyPositionType(slotId)
	//参数不对
	if !slotPosition.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"slotId":   slotId,
			}).Warn("goldequip:处理脱下装备宝石,参数不对")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	order := csGoldEquipTakeOffGem.GetOrder()
	if order < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"order":    order,
			}).Warn("goldequip:处理脱下装备宝石,参数不对")
	}

	err = takeOffGem(tpl, slotPosition, order)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("goldequip:处理脱下装备宝石,错误")
		return err
	}

	log.Debug("goldequip:处理脱下装备宝石,完成")
	return nil
}

//脱下
func takeOffGem(pl player.Player, pos inventorytypes.BodyPositionType, order int32) (err error) {
	flag := takeOffGemInternal(pl, pos, order)
	if !flag {
		return
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	goldequiplogic.GoldEquipPropertyChanged(pl)

	//脱下成功
	scMsg := pbutil.BuildSCGoldEquipTakeOffGem(int32(pos), order)
	pl.SendMsg(scMsg)

	return nil
}

func takeOffGemInternal(pl player.Player, pos inventorytypes.BodyPositionType, order int32) (flag bool) {
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldequipBag := goldequipManager.GetGoldEquipBag()
	it := goldequipBag.GetByPosition(pos)
	//没有东西
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("goldequip:处理脱下装备宝石,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipGemCanNotTakeOff)
		return
	}
	if !goldequipBag.IfEmbedGem(pos, order) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("goldequip:处理脱下装备宝石,已经被卸下")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipGemAlreadyTakeOff)
		return
	}

	//没有足够的空间放脱下来的宝石
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	num := int32(1)
	if !inventoryManager.HasEnoughSlot(it.GetItemId(), num) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("goldequip:处理脱下装备宝石,背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	itemId := goldequipBag.TakeOffGem(pos, order)
	if itemId == 0 {
		panic(fmt.Errorf("goldequip:take off should more than 0"))
	}

	//添加物品
	reasonText := commonlog.InventoryLogReasonTakeOffGem.String()
	flag = inventoryManager.AddItem(itemId, num, commonlog.InventoryLogReasonTakeOffGem, reasonText)
	if !flag {
		panic(fmt.Errorf("goldequip:add item should be success"))
	}
	return
}
