package listener

import (
	"fgame/fgame/core/event"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/longgong/pbutil"
	longgongtypes "fgame/fgame/game/longgong/types"
	npceventtypes "fgame/fgame/game/npc/event/types"

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

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeLongGong {
		return
	}

	bioType := n.GetBiologyTemplate().GetBiologyScriptType()
	if bioType != scenetypes.BiologyScriptTypeLongGongBoss {
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
	rankVal := attackPlayer.GetActivityRankValue(activitytypes.ActivityTypeLongGong, longgongtypes.LongGongSceneRankTypeDamage)
	totalValue := rankVal + damage
	attackPlayer.UpdateActivityRankValue(activitytypes.ActivityTypeLongGong, longgongtypes.LongGongSceneRankTypeDamage, totalValue)

	splM := s.GetAllPlayers()
	for _, spl := range splM {
		scMsg := pbutil.BuildSCLonggongSceneBossHpBroadcast(newHp)
		spl.SendMsg(scMsg)
	}
	return nil
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCHPChanged, event.EventListenerFunc(monsterHPChanged))
}
