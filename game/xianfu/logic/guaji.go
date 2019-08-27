package logic

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	xianfuplayer "fgame/fgame/game/xianfu/player"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
)

func CheckIfPlayerCanEnterXianFu(pl player.Player, xianfuType xianfutypes.XianfuType) bool {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeXinaFu) {
		return false
	}
	xianfuManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	xianfuId := xianfuManager.GetXianfuId(xianfuType)
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xianfuId, xianfuType)
	if xfTemplate == nil {
		return false
	}

	//是否免费次数
	freeTimes := FreeTimesCount(pl, xianfuType)
	if freeTimes < 1 {
		//挑战次数是否足够
		leftTimes := xianfuManager.GetChallengeTimes(xianfuType)
		if leftTimes < 1 {
			return false
		}

		attendNeedItemId := xfTemplate.GetNeedItemId()
		attendNeedItemNum := xfTemplate.GetNeedItemCount()

		//挑战所需物品是否足够
		if !inventoryManager.HasEnoughItem(attendNeedItemId, attendNeedItemNum) {
			return false
		}
	}
	return true
}
