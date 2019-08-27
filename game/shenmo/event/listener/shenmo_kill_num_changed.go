package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playershenmo "fgame/fgame/game/shenmo/player"
)

//玩家击杀人数变更
func playerShenMoGongKillNumChanged(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	manager := p.GetPlayerDataManager(playertypes.PlayerShenMoWarDataManagerType).(*playershenmo.PlayerShenMoDataManager)
	manager.Save()
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShenMoKillNumChanged, event.EventListenerFunc(playerShenMoGongKillNumChanged))
}
