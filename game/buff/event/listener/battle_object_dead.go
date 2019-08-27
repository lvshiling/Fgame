package listener

import (
	"fgame/fgame/core/event"
	bufflogic "fgame/fgame/game/buff/logic"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//死亡
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.BattleObject)

	//获得伤害触发
	bufflogic.TouchBuffByAction(bo, scenetypes.BuffTouchTypeDead)

	//死亡
	bufflogic.RemoveBuffByAction(bo, scenetypes.BuffRemoveTypeDead)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
