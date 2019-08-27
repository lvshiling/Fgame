package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playershenmo "fgame/fgame/game/shenmo/player"
)

//玩家功勋值改变
func battlePlayerShenMoGongXunNumChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerShenMoWarDataManagerType).(*playershenmo.PlayerShenMoDataManager)
	manager.Save()
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShenMoGongXunNumChanged, event.EventListenerFunc(battlePlayerShenMoGongXunNumChanged))
}
