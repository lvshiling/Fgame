package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	playerpk "fgame/fgame/game/pk/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家红名值变化
func playerPkValueChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	pkManager := p.GetPlayerDataManager(playertypes.PlayerPkDataManagerType).(*playerpk.PlayerPkDataManager)
	pkManager.Save()

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerPkValueChanged, event.EventListenerFunc(playerPkValueChanged))
}
