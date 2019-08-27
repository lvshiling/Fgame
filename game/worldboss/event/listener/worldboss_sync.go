package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	playerworldboss "fgame/fgame/game/worldboss/player"
)

//进入场景
func bossReliveDataSync(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	reliveData, ok := data.(*scene.PlayerBossReliveData)
	if !ok {
		return
	}
	worldbossManager := p.GetPlayerDataManager(playertypes.PlayerWorldbossManagerType).(*playerworldboss.PlayerWorldbossManager)
	worldbossManager.BossSync(reliveData.GetBossType(), reliveData.GetReliveTime())
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerBossReliveDataSync, event.EventListenerFunc(bossReliveDataSync))
}
