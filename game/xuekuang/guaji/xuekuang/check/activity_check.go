package check

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

func init() {

	activityguaji.RegisterActivityCheckGuaJi(activitytypes.ActivityTypeXueKuangCollect, activityguaji.ActivityCheckGuaJiFunc(xueKuangCollect))
}

func xueKuangCollect(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) bool {
	return false
}
