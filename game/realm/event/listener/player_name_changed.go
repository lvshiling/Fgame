package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
	playerrealm "fgame/fgame/game/realm/player"
	"fgame/fgame/game/realm/realm"
)

//玩家名字变化
func playerNameChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	realm.GetRealmRankService().PlayerNameChanged(pl)
	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	manager.PlayerChangeName(pl)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerNameChanged, event.EventListenerFunc(playerNameChanged))
}
