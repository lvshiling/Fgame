package listener

import (
	"fgame/fgame/core/event"
	arenapvpeventtypes "fgame/fgame/game/arenapvp/event/types"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//竞技场复活次数变化
func arenapvpReliveChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpObj := arenapvpManager.GetPlayerArenapvpObj()
	reliveTimes := arenapvpObj.GetReliveTimes()
	pl.SetArenapvpReliveTimes(reliveTimes)

	return
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpReliveChanged, event.EventListenerFunc(arenapvpReliveChanged))
}
