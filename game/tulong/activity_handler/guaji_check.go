package activity_handler

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	activityguaji.RegisterActivityCheckGuaJi(activitytypes.ActivityTypeCoressTuLong, activityguaji.ActivityCheckGuaJiFunc(tulongGuaJiCheck))
}

func tulongGuaJiCheck(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) bool {
	return false
	// allianceId := pl.GetAllianceId()
	// if allianceId == 0 {
	// 	return false
	// }

	// if !activitylogic.IfActivityTime(activitytypes.ActivityTypeCoressTuLong)  {
	// 	return false
	// }

	// return true
}
