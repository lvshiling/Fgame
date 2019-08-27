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
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeShield, guaji.GuaJiAdvanceCheckHandlerFunc(shieldCheck))
}

func shieldCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShield) {
		return
	}

	bodyshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	for {
		bodyshieldInfo := bodyshieldManager.GetBodyShiedInfo()
		currentAdvanceId := bodyshieldInfo.ShieldId
		if currentAdvanceId >= advanceId {
			break
		}
		nextAdvancedId := currentAdvanceId + 1
		shieldTemplate := bodyshield.GetBodyShieldService().GetShield(int32(nextAdvancedId))
		if shieldTemplate == nil {
			return
		}

		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		//进阶需要消耗的元宝
		costGold := int64(shieldTemplate.UseGold)
		//进阶需要消耗的银两
		costSilver := int64(shieldTemplate.UseSilver)
		//进阶需要的消耗的绑元
		costBindGold := int64(shieldTemplate.UseBindGold)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := shieldTemplate.UseItem

		useItemTemplate := shieldTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = shieldTemplate.ItemCount
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

		bodyshieldlogic.HandleShieldAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiShieldAdvanced)
	}
}
