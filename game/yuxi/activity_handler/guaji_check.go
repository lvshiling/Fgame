package activity_handler

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	activityguaji.RegisterActivityCheckGuaJi(activitytypes.ActivityTypeYuXi, activityguaji.ActivityCheckGuaJiFunc(yuXiGuaJiCheck))
}

func yuXiGuaJiCheck(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) bool {
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		return false
	}
	return true
}
