package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	moonlovelogic "fgame/fgame/game/moonlove/logic"

	activitytypes "fgame/fgame/game/activity/types"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeMoonLove, activity.ActivityAttendHandlerFunc(moonlovelogic.PlayerEnterMoonloveArgs))
}
