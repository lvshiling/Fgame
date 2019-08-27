package guaji_check

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeSoulRuins, guaji.GuaJiEnterCheckHandlerFunc(soulruinsEnterCheck))
}

func soulruinsEnterCheck(pl player.Player) bool {
	return false
}
