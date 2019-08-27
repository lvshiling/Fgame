package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//自动复活
func playerAutoReborn(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)

	//TODO 机器人
	//断开
	flag := scenelogic.AutoReborn(pl)
	if !flag {
		pl.Close(nil)
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerAutoReborn, event.EventListenerFunc(playerAutoReborn))
}
