package advance

import (
	"fgame/fgame/common/lang"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	winglogic "fgame/fgame/game/wing/logic"
	playerwing "fgame/fgame/game/wing/player"
	"fgame/fgame/game/wing/wing"
)

func init() {
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeFeather, guaji.GuaJiAdvanceCheckHandlerFunc(featherCheck))
}

func featherCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeFeather) {
		return
	}

	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	for {
		currentAdvanceId := wingManager.GetWingInfo().FeatherId
		if currentAdvanceId >= advanceId {
			break
		}
		nextAdvancedId := currentAdvanceId + 1
		featherTemplate := wing.GetWingService().GetFeather(int32(nextAdvancedId))
		if featherTemplate == nil {
			return
		}

		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		//进阶需要消耗的元宝
		costGold := int64(featherTemplate.UseGold)
		//进阶需要消耗的银两
		costSilver := int64(featherTemplate.UseSilver)
		//进阶需要的消耗的绑元
		costBindGold := int64(featherTemplate.UseBindGold)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := featherTemplate.UseItem

		useItemTemplate := featherTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = featherTemplate.ItemCount
			totalNum = inventoryManager.NumOfItems(int32(useItem))
		}

		if totalNum < itemCount {
			if !autobuy {
				return
			}
			//自动进阶
			needBuyNum := itemCount - totalNum

			if !shop.GetShopService().ShopIsSellItem(useItem) {
				return
			}

			isEnoughBuyTimes, shopIdMap := shoplogic.GetPlayerShopCost(pl, useItem, needBuyNum)
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

		winglogic.HandleFeatherAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiFeatherAdvanced)
	}
}
