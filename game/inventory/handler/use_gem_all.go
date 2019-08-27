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
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_USE_GEM_ALL_TYPE), dispatch.HandlerFunc(handleInventoryUseGemAll))
}

//一键使用装备宝石
func handleInventoryUseGemAll(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理一键使用装备宝石")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSInventoryUseGemAll)
	useGemList := csMsg.GetUseGemList()

	err = useGemAll(tpl, useGemList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理一键使用装备宝石,错误")
		return err
	}
	log.Debug("inventory:处理一键使用装备宝石,完成")
	return nil
}

//一键使用装备宝石
func useGemAll(pl player.Player, useGemList []*uipb.UseGem) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipmentBag := inventoryManager.GetEquipmentBag()
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	realmManager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)

	for _, useGem := range useGemList {
		slotId := useGem.GetSlotId()
		indexList := useGem.GetIndexList()
		orderList := useGem.GetOrderList()
		if len(indexList) != len(orderList) {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"slotId":    slotId,
					"indexList": indexList,
					"orderList": orderList,
				}).Warn("inventory:一键使用装备宝石,镶嵌数和宝石数不一致")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
			return
		}

		slotPosition := inventorytypes.BodyPositionType(slotId)
		if !slotPosition.Valid() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"slotId":   slotId,
				}).Warn("inventory:一键使用装备宝石错误,装备槽位错误")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
			return
		}

		// 一键装备宝石
		for i := 0; i < len(orderList); i++ {
			index := indexList[i]
			order := orderList[i]

			it := inventoryManager.FindItemByIndex(inventorytypes.BagTypeGem, index)
			//物品不存在
			if it == nil || it.IsEmpty() {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"index":    index,
					}).Warn("inventory:一键使用装备宝石,物品不存在")
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
					}).Warn("inventory:一键使用装备宝石,模板不存在")
				playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
				return
			}

			//等级不够
			if gemSlotTemplate.NeedLevel > propertyManager.GetLevel() {
				log.WithFields(
					log.Fields{
						"playerId":     pl.GetId(),
						"slotPosition": slotPosition.String(),
						"order":        order,
					}).Warn("inventory:一键使用装备宝石,等级不够")
				playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
				return
			}

			//天截塔等级不够
			if gemSlotTemplate.NeedLayer > realmManager.GetTianJieTaLevel() {
				log.WithFields(
					log.Fields{
						"playerId":     pl.GetId(),
						"slotPosition": slotPosition.String(),
						"order":        order,
					}).Warn("inventory:一键使用装备宝石,天截塔等级不够")
				playerlogic.SendSystemMessage(pl, lang.RealmLevelTooLow)
				return
			}

			//判断是否有装备
			equip := equipmentBag.GetByPosition(slotPosition)
			if equip == nil || equip.IsEmpty() {
				log.WithFields(
					log.Fields{
						"playerId":     pl.GetId(),
						"slotPosition": slotPosition.String(),
						"order":        order,
					}).Warn("inventory:一键使用装备宝石,装备未装上")
				playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotNoEquip)
				return
			}

			//判断是否已经装备上宝石了
			flag := equipmentBag.IfEmbedGem(slotPosition, order)
			if flag {
				flag = takeOffGemInternal(pl, slotPosition, order)
				if !flag {
					continue
				}
			}

			itemId := it.ItemId
			flag = equipmentBag.PutOnGem(slotPosition, order, itemId)
			if !flag {
				//发送错误信息
				panic(fmt.Errorf("inventory:镶嵌宝石,物品[%d],位置[%s],槽位[%d]应该是可以的", itemId, slotPosition.String(), order))
			}

			//移除物品
			reasonText := commonlog.InventoryLogReasonPutOnGem.String()
			flag, _ = inventoryManager.RemoveIndex(inventorytypes.BagTypeGem, index, 1, commonlog.InventoryLogReasonPutOnGem, reasonText)
			if !flag {
				panic(fmt.Errorf("inventory:卸下宝石,位置[%s],槽位[%d]应该是可以的", slotPosition.String(), order))
			}
		}
	}

	//同步改变
	inventorylogic.UpdateEquipmentProperty(pl)
	inventorylogic.SnapInventoryEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCInventoryUseGemAll()
	pl.SendMsg(scMsg)
	return nil
}
