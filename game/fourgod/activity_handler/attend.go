package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	fourgodlogic "fgame/fgame/game/fourgod/logic"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeFourGod, activity.ActivityAttendHandlerFunc(fourgodlogic.PlayerEnterFourGodSceneArgs))
}
