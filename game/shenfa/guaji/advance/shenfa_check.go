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
	shenfalogic "fgame/fgame/game/shenfa/logic"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
)

func init() {
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeShenfa, guaji.GuaJiAdvanceCheckHandlerFunc(shenfaCheck))
}

func shenfaCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenfa) {
		return
	}

	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	for {
		currentAdvanceId := shenfaManager.GetShenfaAdvanced()
		if currentAdvanceId >= advanceId {
			break
		}
		nextAdvancedId := currentAdvanceId + 1
		shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(int32(nextAdvancedId))
		if shenfaTemplate == nil {
			return
		}

		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		//进阶需要消耗的元宝
		costGold := int64(shenfaTemplate.UseMoney)
		//进阶需要消耗的银两
		costSilver := int64(shenfaTemplate.UseYinliang)
		//进阶需要的消耗的绑元
		costBindGold := int64(0)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := shenfaTemplate.UseItem

		useItemTemplate := shenfaTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = shenfaTemplate.ItemCount
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

		shenfalogic.HandleShenfaAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiShenfaAdvanced)
	}
}
