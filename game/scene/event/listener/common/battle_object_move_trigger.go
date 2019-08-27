package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家复活
func battleObjectMoveTrigger(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	// s := bo.GetScene()
	eventData := data.(*battleeventtypes.BattleObjectMoveTriggerEventData)
	//判断是否是玩家
	destPos := eventData.GetDestPos()
	angle := eventData.GetAngle()
	speed := eventData.GetSpeed()
	scenelogic.Move(bo, destPos, angle, speed, scenetypes.MoveTypeNormal, false, false)

	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattleObjectMoveTrigger, event.EventListenerFunc(battleObjectMoveTrigger))
}
