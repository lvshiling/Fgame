package activity_handler

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	activityguaji.RegisterActivityCheckGuaJi(activitytypes.ActivityTypeGodSiegeQiLin, activityguaji.ActivityCheckGuaJiFunc(godsiegeGuaJiCheck))
	activityguaji.RegisterActivityCheckGuaJi(activitytypes.ActivityTypeGodSiegeHuoFeng, activityguaji.ActivityCheckGuaJiFunc(godsiegeGuaJiCheck))
	activityguaji.RegisterActivityCheckGuaJi(activitytypes.ActivityTypeGodSiegeDuLong, activityguaji.ActivityCheckGuaJiFunc(godsiegeGuaJiCheck))
	activityguaji.RegisterActivityCheckGuaJi(activitytypes.ActivityTypeLocalGodSiegeQiLin, activityguaji.ActivityCheckGuaJiFunc(godsiegeGuaJiCheck))
}

func godsiegeGuaJiCheck(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) bool {
	return true
}
