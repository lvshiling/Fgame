package listener

import (
	"fgame/fgame/core/event"
	playeractivity "fgame/fgame/game/activity/player"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	crosstypes "fgame/fgame/game/cross/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//退出跨服
func battlePlayerExitCross(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	crossType, ok := data.(crosstypes.CrossType)
	if !ok {
		return
	}
	activityType, ok := crossType.CrossTypeToActivityType()
	if !ok {
		return
	}
	//退出活动
	//保存数据
	activityManager := p.GetPlayerDataManager(types.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	activityManager.ExitActivity(activityType)

	return
}

func init() {
	gameevent.AddEventListener(crosseventtypes.EventTypePlayerCrossExit, event.EventListenerFunc(battlePlayerExitCross))
}
