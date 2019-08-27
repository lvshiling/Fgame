package advance

import (
	"fgame/fgame/common/lang"
	fabaologic "fgame/fgame/game/fabao/logic"
	playerfabao "fgame/fgame/game/fabao/player"
	fabaotemplate "fgame/fgame/game/fabao/template"
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
)

func init() {
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeFabao, guaji.GuaJiAdvanceCheckHandlerFunc(fabaoCheck))
}

func fabaoCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeFaBao) {
		return
	}

	fabaoManager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	for {
		fabaoInfo := fabaoManager.GetFaBaoInfo()
		if int32(fabaoInfo.GetAdvancedId()) >= advanceId {
			break
		}
		nextAdvancedId := fabaoInfo.GetAdvancedId() + 1
		fabaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(nextAdvancedId))
		if fabaoTemplate == nil {
			return
		}

		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		//进阶需要消耗的元宝
		costGold := int64(fabaoTemplate.UseMoney)
		//进阶需要消耗的银两
		costSilver := int64(fabaoTemplate.UseYinliang)
		//进阶需要的消耗的绑元
		costBindGold := int64(0)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := fabaoTemplate.UseItem

		useItemTemplate := fabaoTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = fabaoTemplate.ItemCount
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

		fabaologic.HandleFaBaoAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiFabaoAdvanced)
	}
}
