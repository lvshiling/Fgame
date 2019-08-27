package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_USE_GEM_ALL_TYPE), dispatch.HandlerFunc(handleGoldEquipUseGemAll))
}

//一键使用装备宝石
func handleGoldEquipUseGemAll(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理一键使用装备宝石")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGoldEquipUseGemAll)
	useGemList := csMsg.GetUseGemList()

	err = useGemAll(tpl, useGemList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("goldequip:处理一键使用装备宝石,错误")
		return err
	}
	log.Debug("goldequip:处理一键使用装备宝石,完成")
	return nil
}

//一键使用装备宝石
func useGemAll(pl player.Player, useGemList []*uipb.UseGem) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldequipBag := goldequipManager.GetGoldEquipBag()

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
				}).Warn("goldequip:一键使用装备宝石,镶嵌数和宝石数不一致")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
			return
		}

		slotPosition := inventorytypes.BodyPositionType(slotId)
		if !slotPosition.Valid() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"slotId":   slotId,
				}).Warn("goldequip:一键使用装备宝石错误,装备槽位错误")
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
					}).Warn("goldequip:一键使用装备宝石,物品不存在")
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

			//判断等级和境界
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
						"playerId":     pl.GetId(),
						"slotPosition": slotPosition.String(),
						"order":        order,
					}).Warn("goldequip:一键使用装备宝石,装备未装上")
				playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentSlotNoEquip)
				return
			}

			//判断是否已经装备上宝石了
			flag := goldequipBag.IfEmbedGem(slotPosition, order)
			if flag {
				flag = takeOffGemInternal(pl, slotPosition, order)
				if !flag {
					continue
				}
			}

			itemId := it.ItemId
			flag = goldequipBag.PutOnGem(slotPosition, order, itemId)
			if !flag {
				//发送错误信息
				panic(fmt.Errorf("goldequip:镶嵌宝石,物品[%d],位置[%s],槽位[%d]应该是可以的", itemId, slotPosition.String(), order))
			}

			//移除物品
			useItemReason := commonlog.InventoryLogReasonPutOnGem
			flag, _ = inventoryManager.RemoveIndex(inventorytypes.BagTypeGem, index, 1, useItemReason, useItemReason.String())
			if !flag {
				panic(fmt.Errorf("goldequip:卸下宝石,位置[%s],槽位[%d]应该是可以的", slotPosition.String(), order))
			}
		}
	}

	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipUseGemAll, pl, nil)

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	goldequiplogic.GoldEquipPropertyChanged(pl)

	scMsg := pbutil.BuildSCGoldEquipUseGemAll()
	pl.SendMsg(scMsg)
	return nil
}
