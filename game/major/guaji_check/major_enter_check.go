package guaji_check

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeShuangXiu, guaji.GuaJiEnterCheckHandlerFunc(shuangXiuEnterCheck))
}

func shuangXiuEnterCheck(pl player.Player) bool {
	return false
}
