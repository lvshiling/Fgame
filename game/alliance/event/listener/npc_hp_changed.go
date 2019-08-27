package listener

import (
	"fgame/fgame/core/event"
	alliancebossscene "fgame/fgame/game/alliance/boss_scene"
	"fgame/fgame/game/alliance/pbutil"
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
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeAllianceBoss {
		return
	}
	sceneData := s.SceneDelegate().(alliancebossscene.AllianceBossSceneData)
	sceneData.UpdateDamage(attackId, damage)

	level := sceneData.GetLevel()
	scAllianceBossChanged := pbutil.BuildSCAllianceBossChanged(level, n)
	s.BroadcastMsg(scAllianceBossChanged)
	return nil
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCHPChanged, event.EventListenerFunc(npcHPChanged))
}
