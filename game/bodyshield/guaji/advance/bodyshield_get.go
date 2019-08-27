package advance

import (
	playerbodyshield "fgame/fgame/game/bodyshield/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

func init() {
	guaji.RegisterGuaJiAdvanceGetHandler(guajitypes.GuaJiAdvanceTypeBodyshield, guaji.GuaJiAdvanceGetHandlerFunc(bodyshieldGet))
}

func bodyshieldGet(pl player.Player, typ guajitypes.GuaJiAdvanceType) int32 {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeBodyShield) {
		return 0
	}
	bodyshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	bodyshieldInfo := bodyshieldManager.GetBodyShiedInfo()
	return int32(bodyshieldInfo.AdvanceId)
}
