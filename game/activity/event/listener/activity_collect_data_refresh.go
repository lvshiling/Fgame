package listener

import (
	"fgame/fgame/core/event"
	playeractivity "fgame/fgame/game/activity/player"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

//采集数据变化
func battlePlayerActivityCollectDataRefresh(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	collectData, ok := data.(*scene.PlayerActvitiyCollectData)
	if !ok {
		return
	}
	activityType := collectData.GetActivityType()
	countMap := collectData.GetCountMap()
	endTime := collectData.GetEndTime()
	activityManager := pl.GetPlayerDataManager(types.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	activityManager.UpdateActivityCollectData(activityType, countMap, endTime)
 
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityCollectDataRefresh, event.EventListenerFunc(battlePlayerActivityCollectDataRefresh))
}
