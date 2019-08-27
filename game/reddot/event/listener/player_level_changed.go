package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	reddotlogic "fgame/fgame/game/reddot/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

var (
	levelChangedMap = map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]int32{
		welfaretypes.OpenActivityTypeInvest: map[welfaretypes.OpenActivitySubType]int32{
			welfaretypes.OpenActivityInvestSubTypeNewLevel: 1,
		},
		welfaretypes.OpenActivityTypeWelfare: map[welfaretypes.OpenActivitySubType]int32{
			welfaretypes.OpenActivityWelfareSubTypeUpLevel: 1,
		},
	}
)

//玩家等级变化
func playerLevelChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	//检查等级红点
	reddotlogic.CheckReddotByType(pl, levelChangedMap)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerLevelChanged, event.EventListenerFunc(playerLevelChanged))
}
