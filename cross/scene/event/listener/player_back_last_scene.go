package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"

	playerlogic "fgame/fgame/cross/player/logic"
	"fgame/fgame/cross/player/player"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	robotlogic "fgame/fgame/game/robot/logic"
)

func playerBackLastScene(target event.EventTarget, data event.EventData) (err error) {

	switch pl := target.(type) {
	case *player.Player:
		playerlogic.ExitCross(pl)
		break
	case scene.RobotPlayer:
		robotlogic.RemoveRobot(pl)
		break
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerBackLastScene, event.EventListenerFunc(playerBackLastScene))
}
