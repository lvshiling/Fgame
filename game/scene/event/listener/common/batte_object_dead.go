package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//生物死亡
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	if bo.GetScene() == nil {
		return
	}
	attackId := data.(int64)
	scenelogic.Dead(bo, attackId)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
