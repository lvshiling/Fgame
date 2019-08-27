package check

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	xuechilogic "fgame/fgame/game/xuechi/logic"
)

func init() {
	guaji.RegisterGuaJiCheckHandler(guajitypes.GuaJiCheckTypeXueChi, guaji.GuaJiCheckHandlerFunc(xueChiGuaJiCheck))
}

func xueChiGuaJiCheck(pl player.Player) {

	//假如血池的量少于
	if pl.GetBlood() < pl.GetMaxHP() {
		//TODO 使用背包的血药

		//补充血池
		itemId, flag := xuechilogic.FindBestXueYao(pl)
		if !flag {
			return
		}
		xuechilogic.HandleXueChiAutoBuy(pl, itemId, 1)
		playerlogic.SendSystemMessage(pl, lang.GuaJiAutoBuyBlood)
	}
}
