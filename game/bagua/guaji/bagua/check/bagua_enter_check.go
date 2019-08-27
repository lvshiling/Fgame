package check

import (
	bagualogic "fgame/fgame/game/bagua/logic"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeBaGua, guaji.GuaJiEnterCheckHandlerFunc(realmEnterCheck))
}

func realmEnterCheck(pl player.Player) bool {
	return bagualogic.CheckPlayerIfCanEnterBaGua(pl)
}
