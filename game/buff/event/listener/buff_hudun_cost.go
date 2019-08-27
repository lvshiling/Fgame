package listener

import (
	"fgame/fgame/core/event"
	buffeventtypes "fgame/fgame/game/buff/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//护盾
func buffHuDunCost(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.BattleObject)
	eventData := data.(*buffeventtypes.BuffEffectCostEventData)

	if eventData.GetEffectType() != scenetypes.BuffEffectTypeHuDun {
		return
	}

	huDunCost := eventData.GetCostNum()
	scObjectDamage := pbutil.BuildSCObjectDamage(bo, scenetypes.DamageTypeHuDun, huDunCost, 0, 0)
	scenelogic.BroadcastNeighborIncludeSelf(bo, scObjectDamage)
	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffEffectCost, event.EventListenerFunc(buffHuDunCost))
}
