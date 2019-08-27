package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	tulonglogic "fgame/fgame/game/tulong/logic"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeCoressTuLong, activity.ActivityAttendHandlerFunc(tulonglogic.PlayerEnterTuLongScene))
}
