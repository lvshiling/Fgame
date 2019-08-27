package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	"math"
)

const (
	replaceHateRatio = 1.2
)

//仇恨值变化
func hateChanged(target event.EventTarget, data event.EventData) (err error) {
	bo, ok := target.(scene.NPC)
	if !ok {
		return
	}

	//最终目标没死
	if bo.GetForeverAttackTarget() != nil && !bo.GetForeverAttackTarget().IsDead() {
		return
	}
	e, ok := data.(*scene.Enemy)
	if !ok {
		return
	}
	newHate := e.GetHate()
	attackTarget := bo.GetAttackTarget()
	if attackTarget == nil {
		bo.SetAttackTarget(e.BattleObject)
		return
	}
	originEnemy := bo.GetEnemy(attackTarget)
	if originEnemy == nil {
		bo.SetAttackTarget(e.BattleObject)
		return
	}
	currentHate := originEnemy.GetHate()

	if newHate <= int(math.Ceil(float64(currentHate)*replaceHateRatio)) {
		return
	}
	bo.SetAttackTarget(e.BattleObject)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattleObjectHateChanged, event.EventListenerFunc(hateChanged))
}
