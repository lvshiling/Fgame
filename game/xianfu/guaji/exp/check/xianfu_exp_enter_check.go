package check

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	playerguaji "fgame/fgame/game/guaji/player"
	guajitypes "fgame/fgame/game/guaji/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	shoplogic "fgame/fgame/game/shop/logic"
	xianfulogic "fgame/fgame/game/xianfu/logic"
	xianfuplayer "fgame/fgame/game/xianfu/player"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeXianFuExp, guaji.GuaJiEnterCheckHandlerFunc(xianFuExpEnterCheck))
}

func xianFuExpEnterCheck(pl player.Player) bool {

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeXinaFu) {
		return false
	}
	xianfuType := xianfutypes.XianfuTypeExp
	xianfuManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	guaJiManager := pl.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)

	xianfuId := xianfuManager.GetXianfuId(xianfuType)
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xianfuId, xianfuType)
	if xfTemplate == nil {
		return false
	}

	//是否免费次数
	freeTimes := xianfulogic.FreeTimesCount(pl, xianfuType)
	if freeTimes > 0 {
		return true
	}

	//挑战次数是否足够
	leftTimes := xianfuManager.GetChallengeTimes(xianfuType)
	if leftTimes < 1 {
		return false
	}
	//挑战所需物品是否足够
	needItemId := xfTemplate.GetNeedItemId()
	needItemNum := xfTemplate.GetNeedItemCount()
	hasNum := inventoryManager.NumOfItems(needItemId)
	if hasNum > needItemNum {
		return true
	}

	//判断是否要自动购买
	guaJiData := guaJiManager.GetGuaJiType(guajitypes.GuaJiTypeXianFuExp)
	if guaJiData.GetOptionValue(guajitypes.GuaJiTypeExpFuBenOptionTypeAutoBuy) == 0 {
		return false
	}
	needBuyNum := needItemNum - hasNum

	//判断够钱吗
	enoughBuyTimes, shopIdMap := shoplogic.GetGuaJiPlayerShopCost(pl, needItemId, needBuyNum)
	//不够购买次数
	if !enoughBuyTimes {
		return false
	}

	//购买
	for shopId, num := range shopIdMap {
		shoplogic.HandleShopBuy(pl, shopId, num)
	}

	return true
}
