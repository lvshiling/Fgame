package advance

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/bodyshield/bodyshield"
	bodyshieldlogic "fgame/fgame/game/bodyshield/logic"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
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
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeBodyshield, guaji.GuaJiAdvanceCheckHandlerFunc(bodyshieldCheck))
}

func bodyshieldCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeBodyShield) {
		return
	}

	bodyshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	for {
		bodyshieldInfo := bodyshieldManager.GetBodyShiedInfo()
		if int32(bodyshieldInfo.AdvanceId) >= advanceId {
			break
		}
		nextAdvancedId := bodyshieldInfo.AdvanceId + 1
		bodyshieldTemplate := bodyshield.GetBodyShieldService().GetBodyShieldNumber(int32(nextAdvancedId))
		if bodyshieldTemplate == nil {
			return
		}

		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		//进阶需要消耗的元宝
		costGold := int64(bodyshieldTemplate.UseMoney)
		//进阶需要消耗的银两
		costSilver := int64(bodyshieldTemplate.UseYinliang)
		//进阶需要的消耗的绑元
		costBindGold := int64(0)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := bodyshieldTemplate.UseItem

		useItemTemplate := bodyshieldTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = bodyshieldTemplate.ItemCount
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

		bodyshieldlogic.HandleBodyShieldAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiBodyshieldAdvanced)
	}
}
