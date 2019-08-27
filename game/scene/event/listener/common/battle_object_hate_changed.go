package common

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

	bo, ok := target.(scene.Player)
	if !ok {
		return
	}
	e, ok := data.(*scene.Enemy)
	if !ok {
		return
	}

	newHate := e.GetHate()

	attackTarget := bo.GetAttackTarget()
	if attackTarget != nil {
		originEnemy := bo.GetEnemy(attackTarget)
		if originEnemy != nil {
			currentHate := originEnemy.GetHate()
			if newHate <= int(math.Ceil(float64(currentHate)*replaceHateRatio)) {
				return
			}
		}
	}
	//反击人
	battleObj := e.BattleObject
	switch battleObj.(type) {
	case scene.Player:
		bo.SetAttackTarget(e.BattleObject)
		break
	}

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattleObjectHateChanged, event.EventListenerFunc(hateChanged))
}
