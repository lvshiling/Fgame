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
	xiantilogic "fgame/fgame/game/xianti/logic"
	playerxianti "fgame/fgame/game/xianti/player"
	"fgame/fgame/game/xianti/xianti"
)

func init() {
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeXianti, guaji.GuaJiAdvanceCheckHandlerFunc(xiantiCheck))
}

func xiantiCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeXianTiAdvanced) {
		return
	}

	xiantiManager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	for {
		currentAdvanceId := xiantiManager.GetXianTiAdvancedId()
		if currentAdvanceId >= advanceId {
			break
		}
		nextAdvancedId := currentAdvanceId + 1
		xiantiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(nextAdvancedId))
		if xiantiTemplate == nil {
			return
		}

		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		//进阶需要消耗的元宝
		costGold := int64(xiantiTemplate.UseMoney)
		//进阶需要消耗的银两
		costSilver := int64(xiantiTemplate.UseYinliang)
		//进阶需要的消耗的绑元
		costBindGold := int64(0)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := xiantiTemplate.UseItem

		useItemTemplate := xiantiTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = xiantiTemplate.ItemCount
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

		xiantilogic.HandleXianTiAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiXiantiAdvanced)
	}
}
