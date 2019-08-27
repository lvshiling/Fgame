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
	playerproperty "fgame/fgame/game/property/player"
	playerrealm "fgame/fgame/game/realm/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_UNLOCK_GEM_TYPE), dispatch.HandlerFunc(handleGoldEquipUnlockGem))
}

//解锁宝石槽
func handleGoldEquipUnlockGem(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理解锁宝石槽")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGoldEquipUnlockGem)
	slotPosition := inventorytypes.BodyPositionType(csMsg.GetSlotId())
	if !slotPosition.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      slotPosition.String(),
			}).Warn("goldequip:处理解锁宝石槽,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	order := csMsg.GetOrder()
	if order < 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"order":    order,
			}).Warn("goldequip:处理解锁宝石槽,失败")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = unlockGem(tpl, slotPosition, order)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("goldequip:处理解锁宝石槽,错误")

		return err
	}
	log.Debug("goldequip:处理解锁宝石槽,完成")
	return nil
}

//解锁宝石槽
func unlockGem(pl player.Player, slotPosition inventorytypes.BodyPositionType, order int32) (err error) {
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	realmManager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	goldequipBag := goldequipManager.GetGoldEquipBag()

	//是否已经解锁
	if goldequipBag.IsUnlockGem(slotPosition, order) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"slotPosition": slotPosition.String(),
				"order":        order,
			}).Warn("goldequip:处理解锁宝石槽,已经解锁")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipAlreadyUnlock)
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
			}).Warn("goldequip:处理解锁宝石槽,模板不存在")
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
			}).Warn("goldequip:处理解锁宝石槽,等级不够")
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
			}).Warn("goldequip:处理解锁宝石槽,天截塔等级不够")
		playerlogic.SendSystemMessage(pl, lang.RealmLevelTooLow)
		return
	}

	items := gemSlotTemplate.GetNeedItemMap()
	if len(items) != 0 {
		if !inventoryManager.HasEnoughItems(items) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"pos":      slotPosition.String(),
					"order":    order,
				}).Warn("goldequip:处理解锁宝石槽,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	if len(items) != 0 {
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonUnlockGem.String(), slotPosition.String(), order)
		flag := inventoryManager.BatchRemove(items, commonlog.InventoryLogReasonPutOnGem, reasonText)
		if !flag {
			panic(fmt.Errorf("goldequip:处理解锁宝石槽,移除材料应该成功"))
		}
	}

	goldequipBag.UnlockGem(slotPosition, order)

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)

	scMsg := pbutil.BuildSCGoldEquipUnlockGem(int32(slotPosition), order)
	pl.SendMsg(scMsg)
	return nil
}
