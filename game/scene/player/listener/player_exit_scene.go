package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerscene "fgame/fgame/game/scene/player"
)

//玩家玩家进入场景
func playerExitScene(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	sceneManager := p.GetPlayerDataManager(playertypes.PlayerSceneDataManagerType).(*playerscene.PlayerSceneDataManager)
	sceneManager.Save()
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(playerExitScene))
}
