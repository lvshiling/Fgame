package listener

import (
	"fgame/fgame/core/event"
	activityeventtypes "fgame/fgame/game/activity/event/types"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	playerfound "fgame/fgame/game/found/player"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//记录活动参与事件
func activityJoinRecord(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	typ, ok := data.(activitytypes.ActivityType)
	if !ok {
		return
	}

	resType, flag := foundtypes.ActivityTypeToFoundResType(typ)
	if !flag {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	manager.IncreFoundResJoinTimes(resType)
	return
}

func init() {
	gameevent.AddEventListener(activityeventtypes.EventTypeActivityJoin, event.EventListenerFunc(activityJoinRecord))
}
