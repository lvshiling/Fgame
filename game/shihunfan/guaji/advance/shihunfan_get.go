package advance

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playershihunfan "fgame/fgame/game/shihunfan/player"
)

func init() {
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeShihunfan, guaji.GuaJiAdvanceGetHandlerFunc(shihunfanGet))
}

func shihunfanGet(pl player.Player, typ guajitypes.GuaJiAdvanceType) int32 {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShiHunFan) {
		return 0
	}
	shihunfanManager := pl.GetPlayerDataManager(types.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	advanceId := shihunfanManager.GetShiHunFanAdvanced()
	return advanceId
}
