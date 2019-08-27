package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/activity/pbutil"
	"fgame/fgame/game/battle/battle"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//排行数据变化
func battlePlayerActivityRankDataChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*battle.BattlePlayerActivityRankDataChangedEventData)
	if !ok {
		return
	}

	rankType := eventData.GetRankType()
	activityType := eventData.GetRankData().GetActivityType()
	val := eventData.GetRankData().GetRankValue(rankType)

	isMsg := pbutil.BuildISPlayerActivityRankDataChanged(int32(activityType), rankType.GetRankType(), val)
	p.SendMsg(isMsg)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityRankDataChanged, event.EventListenerFunc(battlePlayerActivityRankDataChanged))
}
