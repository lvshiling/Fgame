package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	sceneeventypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//怪物掉落
func monsterDrop(target event.EventTarget, data event.EventData) (err error) {
	monsterDropData := data.(*sceneeventypes.MonsterDropData)
	n := target.(scene.NPC)
	attackId := monsterDropData.GetAttackId()
	s := n.GetScene()
	so := s.GetSceneObject(attackId)
	num := monsterDropData.GetNum()
	if so == nil {
		scenelogic.Drop(n, scenetypes.DropOwnerTypePlayer, 0, num)
		return
	}
	scenelogic.Drop(n, scenetypes.DropOwnerTypePlayer, attackId, num)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventypes.EventTypeMonsterDrop, event.EventListenerFunc(monsterDrop))
}
