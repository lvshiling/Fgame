package listener

import (
	"fgame/fgame/core/event"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	npceventtypes "fgame/fgame/game/npc/event/types"
	shengtantypes "fgame/fgame/game/shengtan/types"

	"fgame/fgame/game/npc/npc"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//怪物血量变化
func monsterHPChanged(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	s := n.GetScene()

	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeAllianceShengTan {
		return
	}
	eventData := data.(*npc.NPCHPChangedEventData)

	attackId := eventData.GetAttackId()
	if attackId == 0 {
		return
	}
	attackObject := s.GetSceneObject(attackId)
	if attackObject == nil {
		return
	}

	if attackObject.GetSceneObjectType() != scenetypes.BiologyTypePlayer {
		return
	}

	oldHp := eventData.GetOldHP()
	newHp := eventData.GetNewHP()
	damage := oldHp - newHp
	if damage <= 0 {
		return
	}
	attackPlayer, ok := attackObject.(scene.Player)
	if !ok {
		return
	}
	rankVal := attackPlayer.GetActivityRankValue(activitytypes.ActivityTypeAllianceShengTan, shengtantypes.ShengTanSceneRankTypeDamage)
	totalValue := rankVal + damage
	attackPlayer.UpdateActivityRankValue(activitytypes.ActivityTypeAllianceShengTan, shengtantypes.ShengTanSceneRankTypeDamage, totalValue)
	return nil
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCHPChanged, event.EventListenerFunc(monsterHPChanged))
}
