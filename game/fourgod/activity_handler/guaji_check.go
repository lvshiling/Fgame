package activity_handler

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	activityguaji.RegisterActivityCheckGuaJi(activitytypes.ActivityTypeFourGod, activityguaji.ActivityCheckGuaJiFunc(fourGodGuaJiCheck))
}

func fourGodGuaJiCheck(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) bool {
	return true
}
