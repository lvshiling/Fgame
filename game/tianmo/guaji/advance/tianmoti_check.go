package advance

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
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
	tianmologic "fgame/fgame/game/tianmo/logic"
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotemplate "fgame/fgame/game/tianmo/template"
)

func init() {
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeTianmoti, guaji.GuaJiAdvanceCheckHandlerFunc(tianmoCheck))
}

func tianmoCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeTianMoAdvanced) {
		return
	}

	tianmoManager := pl.GetPlayerDataManager(types.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	for i := 0; i < 100; i++ {
		currentAdvanceId := tianmoManager.GetTianMoAdvanced()
		if currentAdvanceId >= advanceId {
			break
		}
		nextAdvancedId := currentAdvanceId + 1
		tianmoTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(int32(nextAdvancedId))
		if tianmoTemplate == nil {
			return
		}

		tianmoInfo := tianmoManager.GetTianMoInfo()
		if tianmoTemplate.GetActivateType() == commontypes.SpecialAdvancedTypeCharge {
			if tianmoTemplate.ShengjieValue > int32(tianmoInfo.ChargeVal) {
				return
			}
		} else {
			inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
			propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

			//进阶需要消耗的元宝
			costGold := int64(tianmoTemplate.UseMoney)
			//进阶需要消耗的银两
			costSilver := int64(tianmoTemplate.UseYinliang)
			//进阶需要的消耗的绑元
			costBindGold := int64(0)

			//需要消耗物品
			itemCount := int32(0)
			totalNum := int32(0)
			useItem := tianmoTemplate.UseItem

			useItemTemplate := tianmoTemplate.GetUseItemTemplate()
			if useItemTemplate != nil {
				itemCount = tianmoTemplate.ItemCount
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
		}

		tianmologic.HandleTianMoAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiTianmotiAdvanced)
	}
}
