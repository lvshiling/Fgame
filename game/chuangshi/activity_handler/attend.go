package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeChuangShiZhiZhan, activity.ActivityAttendHandlerFunc(playerEnterActivity))
}

func playerEnterActivity(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) (flag bool, err error) {
	cityId := int64(0)
	cityArg := fmt.Sprintf("%d", cityId)
	//进入跨服
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeChuangShi, cityArg)
	flag = true
	return
}
