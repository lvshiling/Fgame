package advance

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerxianti "fgame/fgame/game/xianti/player"
)

func init() {
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeXianti, guaji.GuaJiAdvanceGetHandlerFunc(xiantiGet))
}

func xiantiGet(pl player.Player, typ guajitypes.GuaJiAdvanceType) int32 {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeXianTi) {
		return 0
	}
	xiantiManager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	return xiantiManager.GetXianTiAdvancedId()
}
