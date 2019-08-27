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
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_OPENLIGHT_BAG_TYPE), dispatch.HandlerFunc(handleGoldEquipOpenLightBag))
}

//处理背包金装开光信息
func handleGoldEquipOpenLightBag(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理背包金装开光信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSGoldEquipOpenLightBag)
	index := csMsg.GetItemIdex()
	autoFlag := csMsg.GetIsAuto()

	err = goldequipOpenLightBag(tpl, index, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"autoFlag": autoFlag,
				"error":    err,
			}).Error("goldequip:处理背包金装开光信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"index":    index,
			"autoFlag": autoFlag,
		}).Debug("goldequip:处理背包金装开光完成")
	return nil
}

//背包金装开光的逻辑
func goldequipOpenLightBag(pl player.Player, targetIndex int32, autoFlag bool) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	targetIt := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, targetIndex)

	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
				"autoFlag":    autoFlag,
			}).Warn("goldequip:金装开光失败,强化目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	targetItemTemplate := item.GetItemService().GetItem(int(targetIt.ItemId))
	if !targetItemTemplate.IsGoldEquip() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
				"itemId":      targetIt.ItemId,
			}).Warn("goldequip:金装开光失败,目标不是元神金装")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	//品质
	if targetItemTemplate.GetQualityType() < itemtypes.ItemQualityTypeOrange {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
				"autoFlag":    autoFlag,
			}).Warn("goldequip:金装开光失败,仅橙色品质元神装备可以进行开光")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipOpenLightNotAllow)
		return
	}

	goldequipTemplate := targetItemTemplate.GetGoldEquipTemplate()
	if goldequipTemplate.GoldeuipOpenlightId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
				"autoFlag":    autoFlag,
			}).Warn("goldequip:金装开光失败，该装备不能开光")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	propertyData := targetIt.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	curOpenLevel := propertyData.OpenLightLevel
	openLightTemplate := goldequipTemplate.GetOpenLightTemplate(curOpenLevel)
	nextOpenTemplate := openLightTemplate.GetNextTemplate()
	if nextOpenTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
				"autoFlag":    autoFlag,
			}).Warn("goldequip:金装开光失败，开光等级已达上限")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipOpenLightFullLevel)
		return
	}

	//开光需要物品
	needItems := nextOpenTemplate.GetNeedItemMap()
	if len(needItems) > 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag && autoFlag == false {
			log.WithFields(
				log.Fields{
					"playerid":    pl.GetId(),
					"targetIndex": targetIndex,
					"needItemMap": needItems,
				}).Warn("goldequip:金装开光失败,道具不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//获取背包物品和需要购买物品
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	items, buyItems := inventoryManager.GetItemsAndNeedBuy(needItems)
	//计算需要元宝等
	if len(buyItems) != 0 {
		bindGold := int32(0)
		gold := int32(0)
		sliver := int64(0)
		isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayerMap(pl, buyItems)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerid":    pl.GetId(),
				"targetIndex": targetIndex,
				"autoFlag":    autoFlag,
			}).Warn("goldequip:购买物品失败,自动开光已停止")
			playerlogic.SendSystemMessage(pl, lang.ShopOpenLightAutoBuyItemFail)
			return
		}

		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
		gold += int32(shopNeedGold)
		bindGold += int32(shopNeedBindGold)
		sliver += shopNeedSilver

		flag := propertyManager.HasEnoughCost(int64(bindGold), int64(gold), sliver)
		if !flag {
			log.WithFields(log.Fields{
				"playerid":    pl.GetId(),
				"targetIndex": targetIndex,
				"autoFlag":    autoFlag,
			}).Warn("goldequip:元宝不足，无法开光")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}

		reasonGold := commonlog.GoldLogReasonGoldEquipAutoBuy
		reasonSliver := commonlog.SilverLogReasonGoldEquipAutoBuy
		flag = propertyManager.Cost(int64(bindGold), int64(gold), reasonGold, reasonGold.String(), sliver, reasonSliver, reasonSliver.String())
		if !flag {
			panic(fmt.Errorf("goldequip: goldequipOpenLight Cost should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗物品
	if len(items) != 0 {
		reason := commonlog.InventoryLogReasonGoldEquipOpenLight
		flag := inventoryManager.BatchRemove(items, reason, reason.String())
		if !flag {
			panic(fmt.Errorf("goldequip: goldequipOpenLight use item should be ok"))
		}
	}

	//金装开光判断
	sucess := goldequiplogic.GoldEquipOpenLight(pl, propertyData.OpenTimes, nextOpenTemplate)
	flag := inventoryManager.GoldEquipOpenLight(targetIndex, sucess)
	if !flag {
		panic(fmt.Errorf("goldequip: goldequipOpenLight should be ok"))
	}

	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCGoldEquipOpenLightBag(sucess, propertyData.OpenLightLevel, targetIndex)
	pl.SendMsg(scMsg)
	return
}
