package check

import (
	"fgame/fgame/game/guaji/guaji"
	playerguaji "fgame/fgame/game/guaji/player"
	guajitypes "fgame/fgame/game/guaji/types"
	playerinventory "fgame/fgame/game/inventory/player"
	playermaterial "fgame/fgame/game/material/player"
	materialtemplate "fgame/fgame/game/material/template"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	shoplogic "fgame/fgame/game/shop/logic"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeMaterial, guaji.GuaJiEnterCheckHandlerFunc(materialEnterCheck))
}

func materialEnterCheck(pl player.Player) bool {
	//刷新数据
	materialManager := pl.GetPlayerDataManager(playertypes.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	guaJiManager := pl.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)

	materialManager.RefreshData()
	for materialType, _ := range materialtypes.GetMaterialMap() {
		if !playerlogic.CheckCanEnterScene(pl) {
			continue
		}

		//TODO 添加功能开启判断
		if !pl.IsFuncOpen(materialType.GetFuncOpenType()) {
			continue
		}

		materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(materialType)
		if materialTemplate == nil {
			continue
		}

		//是否免费次数
		if materialManager.IsFreeTimes(materialType) {
			return true
		}

		//挑战次数是否足够
		if !materialManager.IsEnoughAttendTimes(materialType, 1) {
			continue
		}

		//挑战所需物品是否足够
		needItemId := materialTemplate.NeedItemId
		needItemNum := materialTemplate.NeedItemCount
		hasNum := inventoryManager.NumOfItems(needItemId)
		if hasNum > needItemNum {
			return true
		}

		//判断是否要自动购买
		guaJiData := guaJiManager.GetGuaJiType(guajitypes.GuaJiTypeMaterial)
		if guaJiData.GetOptionValue(guajitypes.GuaJiTypeMaterialFuBenOptionTypeAutoBuy) == 0 {
			continue
		}
		needBuyNum := needItemNum - hasNum
		//TODO: zrc封装
		//判断够钱吗
		enoughBuyTimes, shopIdMap := shoplogic.GetGuaJiPlayerShopCost(pl, needItemId, needBuyNum)
		//不够购买次数
		if !enoughBuyTimes {
			continue
		}

		//购买
		for shopId, num := range shopIdMap {
			shoplogic.HandleShopBuy(pl, shopId, num)
		}

		return true
	}
	return false
}
