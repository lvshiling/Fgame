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
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_USE_GEM_TYPE), dispatch.HandlerFunc(handleGoldEquipUseGem))
}

//使用装备宝石
func handleGoldEquipUseGem(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理使用装备宝石")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGoldEquipUseGem)
	index := csMsg.GetIndex()
	if index < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("goldequip:处理使用装备宝石,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	slotPosition := inventorytypes.BodyPositionType(csMsg.GetSlotId())
	if !slotPosition.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      slotPosition.String(),
			}).Warn("goldequip:处理使用装备宝石,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	order := csMsg.GetOrder()
	if order < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"order":    order,
			}).Warn("goldequip:处理使用装备宝石,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = useGem(tpl, index, slotPosition, order)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("goldequip:处理使用装备宝石,错误")

		return err
	}
	log.Debug("goldequip:处理使用装备宝石,完成")
	return nil
}

//使用宝石
func useGem(pl player.Player, index int32, slotPosition inventorytypes.BodyPositionType, order int32) (err error) {
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	goldequipBag := goldequipManager.GetGoldEquipBag()

	it := inventoryManager.FindItemByIndex(inventorytypes.BagTypeGem, index)
	//物品不存在
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("goldequip:使用装备宝石,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	// 物品类型
	itemTemp := item.GetItemService().GetItem(int(it.ItemId))
	if !itemTemp.IsGem() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("goldequip:使用装备宝石,不是宝石")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	//判断宝石槽是否解锁
	if !goldequiplogic.CheckGemUnlockByUse(pl, slotPosition, order) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"slotPosition": slotPosition.String(),
				"order":        order,
			}).Warn("goldequip:处理解锁宝石槽,已经解锁")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipNoUnlock)
		return
	}

	//判断是否有装备
	equip := goldequipBag.GetByPosition(slotPosition)
	if equip == nil || equip.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      slotPosition.String(),
				"order":    order,
			}).Warn("goldequip:使用装备宝石,装备未装上")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentSlotNoEquip)
		return
	}

	//判断是否已经装备上宝石了
	flag := goldequipBag.IfEmbedGem(slotPosition, order)
	if flag {
		flag = takeOffGemInternal(pl, slotPosition, order)
		if !flag {
			return
		}
	}

	itemId := it.ItemId
	flag = goldequipBag.PutOnGem(slotPosition, order, itemId)
	if !flag {
		//发送错误信息
		panic(fmt.Errorf("goldequip:镶嵌宝石,位置[%s],槽位[%d]应该是可以的", slotPosition.String(), order))
	}

	//移除物品
	reasonText := commonlog.InventoryLogReasonPutOnGem.String()
	flag, _ = inventoryManager.RemoveIndex(inventorytypes.BagTypeGem, index, 1, commonlog.InventoryLogReasonPutOnGem, reasonText)
	if !flag {
		panic(fmt.Errorf("goldequip:镶嵌宝石,位置[%s],槽位[%d]应该是可以的", slotPosition.String(), order))
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	goldequiplogic.GoldEquipPropertyChanged(pl)

	scMsg := pbutil.BuildSCGoldEquipUseGem(index, int32(slotPosition), order)
	scMsg.Index = &index
	pl.SendMsg(scMsg)
	return nil
}
