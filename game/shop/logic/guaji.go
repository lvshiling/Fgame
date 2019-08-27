package logic

import (
	playerguaji "fgame/fgame/game/guaji/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/shop/shop"
	shoptypes "fgame/fgame/game/shop/types"
	"fmt"
)

//购买所需花费
func GetGuaJiPlayerShopCost(pl player.Player, itemId int32, needNum int32) (isEnoughBuyTimes bool, shopIdMap map[int32]int32) {
	if needNum < 0 {
		panic(fmt.Errorf("shoplogic: needNum 应该大于等于0"))
	}
	if !shop.GetShopService().ShopIsSellItem(itemId) {
		return
	}
	//获取需要剩余的银两
	guaJiManager := pl.GetPlayerDataManager(types.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)
	remainSilver := guaJiManager.GetRemainSilver()
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	shopIdMap = make(map[int32]int32)
	leftNum := needNum
	maxSilver := propertyManager.GetSilver() - remainSilver
	if maxSilver > 0 {
		silverMaxNum, silverShopMap, _ := GetPlayerShopCostByComsumeType(pl, maxSilver, shoptypes.ShopConsumeTypeSliver, itemId, needNum)
		if silverMaxNum > 0 {
			leftNum -= silverMaxNum

			for shopId, num := range silverShopMap {
				shopIdMap[shopId] = num
			}
		}
	}

	if leftNum < 0 {
		panic(fmt.Errorf("剩余数量应该不小于0"))
	}

	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}
	maxBindGold := propertyManager.GetBindGlod()
	maxGold := propertyManager.GetGold()

	bindGoldMaxNum, bindGoldShopMap, bindGoldCost := GetPlayerShopCostByComsumeType(pl, maxBindGold+maxGold, shoptypes.ShopConsumeTypeBindGold, itemId, leftNum)
	if bindGoldMaxNum > 0 {
		leftNum -= bindGoldMaxNum

		for shopId, num := range bindGoldShopMap {
			shopIdMap[shopId] = num
		}
	}
	if leftNum < 0 {
		panic(fmt.Errorf("剩余数量应该不小于0"))
	}

	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}
	remainGold := maxGold
	if bindGoldCost > maxBindGold {
		remainGold -= (bindGoldCost - maxBindGold)
	}
	goldMaxNum, goldShopMap, _ := GetPlayerShopCostByComsumeType(pl, remainGold, shoptypes.ShopConsumeTypeGold, itemId, leftNum)
	if goldMaxNum > 0 {
		leftNum -= goldMaxNum

		for shopId, num := range goldShopMap {
			shopIdMap[shopId] = num
		}
	}
	if leftNum < 0 {
		panic(fmt.Errorf("剩余数量应该不小于0"))
	}

	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}
	return
}
