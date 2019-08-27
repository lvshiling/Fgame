package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeXueKuangCollect, activity.ActivityAttendHandlerFunc(PlayerXueKuangEnterScene))
}

func PlayerXueKuangEnterScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeXueKuang)
	flag = true
	return
}
