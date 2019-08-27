package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	npceventtypes "fgame/fgame/game/npc/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//怪物改变
func npcCampChanged(target event.EventTarget, data event.EventData) (err error) {
	n := target.(scene.NPC)
	for _, obj := range n.GetNeighbors() {
		p, ok := obj.(scene.Player)
		if !ok {
			continue
		}
		isEnemy := p.IsEnemy(n)
		scMonsterCampChanged := pbutil.BuildSCMonsterCampChanged(n, isEnemy)
		p.SendMsg(scMonsterCampChanged)
	}
	return
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCCampChanged, event.EventListenerFunc(npcCampChanged))
}
