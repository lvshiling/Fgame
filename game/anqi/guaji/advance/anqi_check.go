package advance

import (
	"fgame/fgame/common/lang"
	anqilogic "fgame/fgame/game/anqi/logic"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
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
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeAnqi, guaji.GuaJiAdvanceCheckHandlerFunc(anqiCheck))
}

func anqiCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeAnQi) {
		return
	}

	anqiManager := pl.GetPlayerDataManager(types.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	for {
		anqiInfo := anqiManager.GetAnqiInfo()
		if int32(anqiInfo.AdvanceId) >= advanceId {
			break
		}
		nextAdvancedId := anqiInfo.AdvanceId + 1
		anqiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(nextAdvancedId))
		if anqiTemplate == nil {
			return
		}

		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		//进阶需要消耗的元宝
		costGold := int64(anqiTemplate.UseMoney)
		//进阶需要消耗的银两
		costSilver := int64(anqiTemplate.UseYinliang)
		//进阶需要的消耗的绑元
		costBindGold := int64(0)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := anqiTemplate.UseItem

		useItemTemplate := anqiTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = anqiTemplate.ItemCount
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

		anqilogic.HandleAnqiAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiAnqiAdvanced)
	}
}
