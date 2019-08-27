package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/worldboss/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//从血池补血
func reliveDataSync(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	playerBossReliveData, ok := data.(*scene.PlayerBossReliveData)
	isPlayerBossReliveSync := pbutil.BuildISPlayerBossReliveSync(pl, playerBossReliveData)
	pl.SendMsg(isPlayerBossReliveSync)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerBossReliveDataSync, event.EventListenerFunc(reliveDataSync))
}
