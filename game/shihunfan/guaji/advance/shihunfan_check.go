package advance

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
)

func init() {
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeShihunfan, guaji.GuaJiAdvanceCheckHandlerFunc(shihunfanCheck))
}

func shihunfanCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShiHunFanAdvanced) {
		return
	}
	//TODO 特殊处理
	// shihunfanManager := pl.GetPlayerDataManager(types.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	// for {
	// 	currentAdvanceId := shihunfanManager.GetShiHunFanAdvanced()
	// 	if currentAdvanceId >= advanceId {
	// 		break
	// 	}
	// 	nextAdvancedId := currentAdvanceId + 1
	// 	shihunfanTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(int32(nextAdvancedId))
	// 	if shihunfanTemplate == nil {
	// 		return
	// 	}

	// 	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	// 	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	// 	//进阶需要消耗的元宝
	// 	costGold := int64(shihunfanTemplate.UseMoney)
	// 	//进阶需要消耗的银两
	// 	costSilver := int64(shihunfanTemplate.UseYinliang)
	// 	//进阶需要的消耗的绑元
	// 	costBindGold := int64(0)

	// 	//需要消耗物品
	// 	itemCount := int32(0)
	// 	totalNum := int32(0)
	// 	useItem := shihunfanTemplate.UseItem

	// 	useItemTemplate := shihunfanTemplate.GetUseItemTemplate()
	// 	if useItemTemplate != nil {
	// 		itemCount = shihunfanTemplate.ItemCount
	// 		totalNum = inventoryManager.NumOfItems(int32(useItem))
	// 	}

	// 	if totalNum < itemCount {
	// 		if !autobuy {
	// 			return
	// 		}
	// 		//自动进阶
	// 		needBuyNum := itemCount - totalNum

	// 		if !shop.GetShopService().ShopIsSellItem(useItem) {
	// 			return
	// 		}

	// 		isEnoughBuyTimes, shopIdMap := shoplogic.GetPlayerShopCost(pl, useItem, needBuyNum)
	// 		if !isEnoughBuyTimes {
	// 			return
	// 		}

	// 		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
	// 		costGold += shopNeedGold
	// 		costBindGold += shopNeedBindGold
	// 		costSilver += shopNeedSilver

	// 	}
	// 	flag := propertyManager.HasEnoughCost(costBindGold, costGold, costSilver)
	// 	if !flag {
	// 		return
	// 	}

	// 	shihunfanlogic.HandleShiHunFanAdvanced(pl, autobuy)
	// 	playerlogic.SendSystemMessage(pl, lang.GuaJiShihunfanAdvanced)
	// }
}
