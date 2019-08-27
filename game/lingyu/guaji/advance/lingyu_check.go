package advance

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	playerinventory "fgame/fgame/game/inventory/player"
	lingyulogic "fgame/fgame/game/lingyu/logic"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
)

func init() {
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeLingyu, guaji.GuaJiAdvanceCheckHandlerFunc(lingyuCheck))
}

func lingyuCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeLingYu) {
		return
	}

	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	//以防循环死
	for i := 0; i <= 100; i++ {
		lingyuInfo := lingyuManager.GetLingyuInfo()
		if int32(lingyuInfo.AdvanceId) >= advanceId {
			break
		}
		nextAdvancedId := lingyuInfo.AdvanceId + 1
		lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(int32(nextAdvancedId))
		if lingyuTemplate == nil {
			return
		}
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		if lingyuTemplate.GetActivateType() == commontypes.SpecialAdvancedTypeCost {
			needActivateGold := int64(lingyuTemplate.ShengjieValue)
			if !propertyManager.HasEnoughGold(needActivateGold, false) {
				return
			}
		} else {
			//进阶需要消耗的元宝
			costGold := int64(lingyuTemplate.UseMoney)
			//进阶需要消耗的银两
			costSilver := int64(lingyuTemplate.UseYinliang)
			//进阶需要的消耗的绑元
			costBindGold := int64(0)

			//需要消耗物品
			itemCount := int32(0)
			totalNum := int32(0)
			useItem := lingyuTemplate.UseItem

			useItemTemplate := lingyuTemplate.GetUseItemTemplate()
			if useItemTemplate != nil {
				itemCount = lingyuTemplate.ItemCount
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
		lingyulogic.HandleLingyuAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiLingyuAdvanced)
	}
}
