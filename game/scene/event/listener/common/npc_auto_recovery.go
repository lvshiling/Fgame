package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	npceventtypes "fgame/fgame/game/npc/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//怪物改变
func npcAutoRecoverChanged(target event.EventTarget, data event.EventData) (err error) {
	n := target.(scene.NPC)
	recoverHp := data.(int64)
	for _, obj := range n.GetNeighbors() {
		p, ok := obj.(scene.Player)
		if !ok {
			continue
		}

		scObjectDamage := pbutil.BuildSCObjectDamage(n, scenetypes.DamageTypeAutoRecovery, recoverHp, 0, 0)
		p.SendMsg(scObjectDamage)
	}
	return
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCAutoRecover, event.EventListenerFunc(npcAutoRecoverChanged))
}
