package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	longgonglogic "fgame/fgame/game/longgong/logic"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeLongGong, activity.ActivityAttendHandlerFunc(longgonglogic.PlayerEnterLongGongScene))
}
