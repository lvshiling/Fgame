package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/player/pbutil"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//boss被杀
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo, ok := target.(scene.BattleObject)
	if !ok {
		return
	}
	npc, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	biologyTemplate := npc.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}
	biologyId := int32(biologyTemplate.TemplateId())
	attackId, ok := data.(int64)
	if !ok {
		return
	}

	pl := npc.GetScene().GetPlayer(attackId)
	if pl == nil {
		return
	}

	isPlayerKillBiology := pbutil.BuildISPlayerKillBiology(biologyId)
	pl.SendMsg(isPlayerKillBiology)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
