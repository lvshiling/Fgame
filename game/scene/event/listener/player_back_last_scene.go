package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"

	battleeventtypes "fgame/fgame/game/battle/event/types"
	robotlogic "fgame/fgame/game/robot/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

func playerBackLastScene(target event.EventTarget, data event.EventData) (err error) {
	switch p := target.(type) {
	case player.Player:
		scenelogic.PlayerBackLastScene(p)
		break
	case scene.RobotPlayer:
		robotlogic.RemoveRobot(p)
		break
	}
	return err
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerBackLastScene, event.EventListenerFunc(playerBackLastScene))
}
