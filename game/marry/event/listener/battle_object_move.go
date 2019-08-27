package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	hunchenpc "fgame/fgame/game/marry/npc/hunche"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//到达目标点
func battleObjectMove(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeWeddingCar {
		return
	}

	hunChe := n.(*hunchenpc.HunCheNPC)
	hunCheObj := hunChe.GetHunCheObject()
	//结婚两人瞬移
	plObj := hunChe.GetScene().GetSceneObject(hunCheObj.GetPlayerId())
	if plObj != nil {
		pl, ok := plObj.(player.Player)
		if ok {
			scenelogic.FixPosition(pl, bo.GetPosition())
		}
	}
	spousePlObj := hunChe.GetScene().GetSceneObject(hunCheObj.GetSpouseId())
	if spousePlObj != nil {
		spousePl, ok := spousePlObj.(player.Player)
		if ok {
			scenelogic.FixPosition(spousePl, bo.GetPosition())
		}
	}
	//判断是否到达
	if !hunCheObj.IsReachGoal(hunChe.GetPosition()) {
		return
	}

	hunChe.ReachGoal()
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectMove, event.EventListenerFunc(battleObjectMove))
}
