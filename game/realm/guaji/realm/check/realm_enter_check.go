package check

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	realmlogic "fgame/fgame/game/realm/logic"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeTianJieTa, guaji.GuaJiEnterCheckHandlerFunc(realmEnterCheck))
}

func realmEnterCheck(pl player.Player) bool {
	return realmlogic.CheckIfCanEnterTianJieTa(pl)
}
