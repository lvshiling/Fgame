package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerfound "fgame/fgame/game/found/player"
	foundtypes "fgame/fgame/game/found/types"
	materialeventtypes "fgame/fgame/game/material/event/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//记录材料副本挑战事件
func materialChallenge(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*materialeventtypes.MaterialChallengeEventData)
	if !ok {
		return
	}
	materialType := eventData.GetType()
	num := eventData.GetNum()
	if num <= 0 {
		return
	}

	resType, flag := foundtypes.MaterialTypeToFoundResType(materialType)
	if !flag {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	manager.IncreFoundResJoinTimesBatch(resType, num)

	return
}

func init() {
	gameevent.AddEventListener(materialeventtypes.EventTypeMaterialChallenge, event.EventListenerFunc(materialChallenge))

}
