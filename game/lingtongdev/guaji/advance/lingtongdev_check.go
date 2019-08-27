package advance

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	playerinventory "fgame/fgame/game/inventory/player"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
)

func init() {
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeLingTongWeapon, guaji.GuaJiAdvanceCheckHandlerFunc(lingTongDevCheck))
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeLingTongMount, guaji.GuaJiAdvanceCheckHandlerFunc(lingTongDevCheck))
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeLingTongWing, guaji.GuaJiAdvanceCheckHandlerFunc(lingTongDevCheck))
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeLingTongShenFa, guaji.GuaJiAdvanceCheckHandlerFunc(lingTongDevCheck))
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeLingTongLingYu, guaji.GuaJiAdvanceCheckHandlerFunc(lingTongDevCheck))
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeLingTongFaBao, guaji.GuaJiAdvanceCheckHandlerFunc(lingTongDevCheck))
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeLingTongXianTi, guaji.GuaJiAdvanceCheckHandlerFunc(lingTongDevCheck))

}

func lingTongDevCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	classType := getLingTongDevType(typ)
	if !classType.Vaild() {
		return
	}
	if !pl.IsFuncOpen(classType.GetAdvanceFuncOpenType()) {
		return
	}
	lingTongManager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongMap := lingTongManager.GetLingTongMap()
	if len(lingTongMap) == 0 {
		return
	}

	lingTongDevManager := pl.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := lingTongDevManager.GetLingTongDevInfo(classType)
	if lingTongDevInfo == nil {
		return
	}
	for {
		currentAdvanceId := lingTongDevInfo.GetAdvancedId()
		if currentAdvanceId >= advanceId {
			break
		}
		nextAdvancedId := currentAdvanceId + 1
		lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, nextAdvancedId)
		if lingTongDevTemplate == nil {
			return
		}

		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		//进阶需要消耗的元宝
		costGold := lingTongDevTemplate.GetGold()
		//进阶需要消耗的银两
		costSilver := lingTongDevTemplate.GetSilver()
		//进阶需要消耗的绑元
		costBindGold := int64(0)

		needItems := lingTongDevTemplate.GetItemMap()
		if len(needItems) != 0 {
			flag := inventoryManager.HasEnoughItems(needItems)
			if !flag && !autobuy {
				return
			}
		}

		isEnoughBuyTimes := true
		shopIdMap := make(map[int32]int32)
		//获取背包物品和需要购买物品
		_, buyItems := inventoryManager.GetItemsAndNeedBuy(needItems)
		//计算需要元宝等
		if len(buyItems) != 0 {
			isEnoughBuyTimes, shopIdMap = shoplogic.GetPlayerShopCostForItemMap(pl, buyItems)
			if !isEnoughBuyTimes {
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			costGold += shopNeedGold
			costBindGold += shopNeedBindGold
			costSilver += shopNeedSilver
		}
		flag := propertyManager.HasEnoughCost(costBindGold, costGold, costSilver)
		if !flag {
			return
		}
		lingtongdevlogic.HandleLingTongDevAdvanced(pl, classType, autobuy)
		playerlogic.SendSystemMessage(pl, getLingTongDevLangCode(typ))
	}
}
