package listener

import (
	"fgame/fgame/core/event"
	playeractivity "fgame/fgame/game/activity/player"
	"fgame/fgame/game/battle/battle"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//排行数据变化
func battlePlayerActivityRankDataChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*battle.BattlePlayerActivityRankDataChangedEventData)
	if !ok {
		return
	}
	activityType := eventData.GetRankData().GetActivityType()

	rankMap := eventData.GetRankData().GetRankMap()
	endTime := eventData.GetRankData().GetEndTime()
	activityManager := p.GetPlayerDataManager(types.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	activityManager.UpdateActivityRankData(activityType, rankMap, endTime)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityRankDataChanged, event.EventListenerFunc(battlePlayerActivityRankDataChanged))
}
