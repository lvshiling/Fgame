package advance

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertianmo "fgame/fgame/game/tianmo/player"
)

func init() {
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeTianmoti, guaji.GuaJiAdvanceGetHandlerFunc(tianmotiGet))
}

func tianmotiGet(pl player.Player, typ guajitypes.GuaJiAdvanceType) int32 {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeTianMo) {
		return 0
	}
	tianmotiManager := pl.GetPlayerDataManager(types.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	advanceId := tianmotiManager.GetTianMoAdvanced()
	return advanceId
}
