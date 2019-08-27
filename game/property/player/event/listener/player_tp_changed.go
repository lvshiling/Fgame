package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
)

//玩家体力变化
func playerTpChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	m := p.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	m.SetTp(p.GetTP())
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerTPChanged, event.EventListenerFunc(playerTpChanged))
}
