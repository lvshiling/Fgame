package listener

import (
	"fgame/fgame/core/event"
	bufflogic "fgame/fgame/game/buff/logic"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//移动
func battleObjectCure(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	//移除被攻击打断
	bufflogic.RemoveBuffByAction(bo, scenetypes.BuffRemoveTypeCure)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectCure, event.EventListenerFunc(battleObjectCure))
}
