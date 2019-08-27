package activity_handler

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	activityguaji.RegisterActivityCheckGuaJi(activitytypes.ActivityTypeAllianceShengTan, activityguaji.ActivityCheckGuaJiFunc(shengTanGuaJiCheck))
}

func shengTanGuaJiCheck(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) bool {
	if pl.GetAllianceId() == 0 {
		return false
	}
	return true
}
