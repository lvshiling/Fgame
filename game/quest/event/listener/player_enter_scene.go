package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	mapType := s.MapTemplate().GetMapType()
	err = enterTower(pl, mapType)
	if err != nil {
		return
	}
	return
}

//进入打宝塔
func enterTower(pl player.Player, mapType scenetypes.SceneType) (err error) {
	return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEnterTower, 0, 1)
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
