package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	xiantaologic "fgame/fgame/game/xiantao/logic"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeXianTaoDaHui, activity.ActivityAttendHandlerFunc(xiantaologic.PlayerEnterXianTaoScene))
}
