package listener

import (
	"fgame/fgame/core/event"
	bufflogic "fgame/fgame/game/buff/logic"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//攻击
func battleObjectAttack(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.BattleObject)

	//获得伤害触发
	bufflogic.TouchBuffByAction(bo, scenetypes.BuffTouchTypeObjectDamageSelf)

	//获得伤害触发
	bufflogic.TouchBuffByAction(bo, scenetypes.BuffTouchTypeObjectDamage)

	bufflogic.RemoveBuffByAction(bo, scenetypes.BuffRemoveTypeAttack)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectAttack, event.EventListenerFunc(battleObjectAttack))
}
