package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerrelive "fgame/fgame/game/relive/player"
)

//加载完成后
func playerReliveRefresh(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	m := p.GetPlayerDataManager(playertypes.PlayerReliveDataManagerType).(*playerrelive.PlayerReliveDataManager)
	m.Save()
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerReliveRefresh, event.EventListenerFunc(playerReliveRefresh))
}
