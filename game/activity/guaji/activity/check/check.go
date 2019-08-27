package check

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeActivity, guaji.GuaJiEnterCheckHandlerFunc(activityEnterCheck))
}

func activityEnterCheck(pl player.Player) bool {
	activityTemplate := activityguaji.GetGuaJiActivityTemplate(pl)
	if activityTemplate == nil {
		return false
	}
	return true
}
