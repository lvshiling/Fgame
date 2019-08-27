package advance

import (
	"fgame/fgame/common/lang"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	playerinventory "fgame/fgame/game/inventory/player"
	massacrelogic "fgame/fgame/game/massacre/logic"
	playermassacre "fgame/fgame/game/massacre/player"
	massacretemplate "fgame/fgame/game/massacre/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
)

func init() {
	guaji.RegisterGuaJiAdvanceCheckHandler(guajitypes.GuaJiAdvanceTypeMassacre, guaji.GuaJiAdvanceCheckHandlerFunc(massacreCheck))
}

func massacreCheck(pl player.Player, typ guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeMassacre) {
		return
	}

	massacreManager := pl.GetPlayerDataManager(types.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	for {
		massacreInfo := massacreManager.GetMassacreInfo()
		currentAdvanceId := int32(massacreInfo.AdvanceId)
		if currentAdvanceId >= advanceId {
			break
		}
		nextAdvancedId := currentAdvanceId + 1

		massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(int(nextAdvancedId))
		if massacreTemplate == nil {
			return
		}

		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

		//需要消耗的杀气
		needSqNum := int64(massacreTemplate.UseGas)

		//进阶需要消耗的银两
		costSilver := int64(massacreTemplate.UseMoney)
		//进阶需要消耗的元宝
		costGold := int64(0)
		//进阶需要消耗的绑元
		costBindGold := int64(0)

		//需要消耗物品
		itemCount := int32(0)
		totalNum := int32(0)
		useItem := massacreTemplate.UseItem
		isEnoughBuyTimes := true
		shopIdMap := make(map[int32]int32)
		useItemTemplate := massacreTemplate.GetUseItemTemplate()
		if useItemTemplate != nil {
			itemCount = massacreTemplate.ItemCount
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

			isEnoughBuyTimes, shopIdMap = shoplogic.GetPlayerShopCost(pl, useItem, needBuyNum)
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

		if needSqNum > 0 && needSqNum > massacreInfo.ShaQiNum {
			return
		}
		massacrelogic.HandleMassacreAdvanced(pl, autobuy)
		playerlogic.SendSystemMessage(pl, lang.GuaJiMassacreAdvanced)
	}
}
