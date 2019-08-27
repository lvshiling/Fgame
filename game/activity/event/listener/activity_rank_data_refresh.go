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

//排行数据变化
func battlePlayerActivityRankDataRefresh(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	rankData, ok := data.(*scene.PlayerActvitiyRankData)
	if !ok {
		return
	}
	activityType := rankData.GetActivityType()
	rankMap := rankData.GetRankMap()
	endTime := rankData.GetEndTime()
	activityManager := p.GetPlayerDataManager(types.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	activityManager.UpdateActivityRankData(activityType, rankMap, endTime)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityRankDataRefresh, event.EventListenerFunc(battlePlayerActivityRankDataRefresh))
}
