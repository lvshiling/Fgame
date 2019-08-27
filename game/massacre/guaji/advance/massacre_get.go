package advance

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	playermassacre "fgame/fgame/game/massacre/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

func init() {
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeMassacre, guaji.GuaJiAdvanceGetHandlerFunc(massacreGet))
}

func massacreGet(pl player.Player, typ guajitypes.GuaJiAdvanceType) int32 {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeMassacre) {
		return 0
	}
	massacreManager := pl.GetPlayerDataManager(types.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	return int32(massacreManager.GetMassacreInfo().AdvanceId)
}
