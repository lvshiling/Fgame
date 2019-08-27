package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerfound "fgame/fgame/game/found/player"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	soulruinseventtypes "fgame/fgame/game/soulruins/event/types"
)

//记录帝魂遗迹接受事件
func soulruinsAcceptRecord(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	num := data.(int32)
	if num <= 0 {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)

	for num > 0 {
		num -= 1
		manager.IncreFoundResJoinTimes(foundtypes.FoundResourceTypeDiHunYiJi)
	}
	return
}

func init() {
	gameevent.AddEventListener(soulruinseventtypes.EventTypeSoulruinsChallenge, event.EventListenerFunc(soulruinsAcceptRecord))
}
