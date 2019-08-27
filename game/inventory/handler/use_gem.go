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
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	playerproperty "fgame/fgame/game/property/player"
	playerrealm "fgame/fgame/game/realm/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_USE_GEM_TYPE), dispatch.HandlerFunc(handleInventoryUseGem))
}

//使用装备宝石
func handleInventoryUseGem(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理使用装备宝石")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryUseGem := msg.(*uipb.CSInventoryUseGem)
	index := csInventoryUseGem.GetIndex()
	if index < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:处理使用装备宝石,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	slotPosition := inventorytypes.BodyPositionType(csInventoryUseGem.GetSlotId())
	if !slotPosition.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      slotPosition.String(),
			}).Warn("inventory:处理使用装备宝石,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	order := csInventoryUseGem.GetOrder()
	if order < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"order":    order,
			}).Warn("inventory:处理使用装备宝石,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = useGem(tpl, index, slotPosition, order)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理使用装备宝石,错误")

		return err
	}
	log.Debug("inventory:处理使用装备宝石,完成")
	return nil
}

//使用宝石
func useGem(pl player.Player, index int32, slotPosition inventorytypes.BodyPositionType, order int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipmentBag := manager.GetEquipmentBag()

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	realmManager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	it := manager.FindItemByIndex(inventorytypes.BagTypeGem, index)
	//物品不存在
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用装备宝石,物品不存在")
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
			}).Warn("inventory:使用装备宝石,不是宝石")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	//判断等级和境界
	gemSlotTemplate := item.GetItemService().GetEquipGemSlotTemplate(slotPosition, order)
	if gemSlotTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"slotPosition": slotPosition.String(),
				"order":        order,
			}).Warn("inventory:使用装备宝石,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	//等级不够
	if gemSlotTemplate.NeedLevel > propertyManager.GetLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      slotPosition.String(),
				"order":    order,
			}).Warn("inventory:使用装备宝石,等级不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//天截塔等级不够
	if gemSlotTemplate.NeedLayer > realmManager.GetTianJieTaLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      slotPosition.String(),
				"order":    order,
			}).Warn("inventory:使用装备宝石,天截塔等级不够")
		playerlogic.SendSystemMessage(pl, lang.RealmLevelTooLow)
		return
	}

	//判断是否有装备
	equip := equipmentBag.GetByPosition(slotPosition)
	if equip == nil || equip.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      slotPosition.String(),
				"order":    order,
			}).Warn("inventory:使用装备宝石,装备未装上")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotNoEquip)
		return
	}

	//判断是否已经装备上宝石了
	flag := equipmentBag.IfEmbedGem(slotPosition, order)
	if flag {
		flag = takeOffGemInternal(pl, slotPosition, order)
		if !flag {
			return
		}
	}

	itemId := it.ItemId
	flag = equipmentBag.PutOnGem(slotPosition, order, itemId)
	if !flag {
		//发送错误信息
		panic(fmt.Errorf("inventory:镶嵌宝石,位置[%s],槽位[%d]应该是可以的", slotPosition.String(), order))
	}

	//移除物品
	reasonText := commonlog.InventoryLogReasonPutOnGem.String()
	flag, _ = manager.RemoveIndex(inventorytypes.BagTypeGem, index, 1, commonlog.InventoryLogReasonPutOnGem, reasonText)
	if !flag {
		panic(fmt.Errorf("inventory:镶嵌宝石,位置[%s],槽位[%d]应该是可以的", slotPosition.String(), order))
	}

	//同步改变
	inventorylogic.UpdateEquipmentProperty(pl)
	logic.SnapInventoryEquipChanged(pl)
	logic.SnapInventoryChanged(pl)

	scInventoryUseEquip := pbutil.BuildSCInventoryUseEquip(index)
	scInventoryUseEquip.Index = &index
	pl.SendMsg(scInventoryUseEquip)
	return nil
}
