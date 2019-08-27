package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	densewatlogic "fgame/fgame/game/densewat/logic"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeDenseWat, activity.ActivityAttendHandlerFunc(densewatlogic.PlayerEnterDenseWatArgs))
}
