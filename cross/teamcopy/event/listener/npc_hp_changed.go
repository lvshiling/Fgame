package listener

import (
	"fgame/fgame/core/event"
	teamcopyscene "fgame/fgame/cross/teamcopy/scene"
	gameevent "fgame/fgame/game/event"
	npceventtypes "fgame/fgame/game/npc/event/types"
	"fgame/fgame/game/npc/npc"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//npc血量变化
func npcHPChanged(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	eventData := data.(*npc.NPCHPChangedEventData)
	//回血
	attackId := eventData.GetAttackId()
	if attackId == 0 {
		return
	}
	oldHp := eventData.GetOldHP()
	newHp := eventData.GetNewHP()
	damage := oldHp - newHp
	if damage <= 0 {
		return
	}

	s := n.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossTeamCopy {
		return
	}
	sceneData := s.SceneDelegate().(teamcopyscene.TeamCopySceneData)
	sceneData.UpdateDamage(attackId, damage)
	return nil
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCHPChanged, event.EventListenerFunc(npcHPChanged))
}
