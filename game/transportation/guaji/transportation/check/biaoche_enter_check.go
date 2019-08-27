package check

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	transportationlogic "fgame/fgame/game/transportation/logic"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeBiaoChe, guaji.GuaJiEnterCheckHandlerFunc(biaoCheEnterCheck))
}

func biaoCheEnterCheck(pl player.Player) bool {
	_, flag := transportationlogic.CheckIfCanTransportaion(pl)
	return flag
}
