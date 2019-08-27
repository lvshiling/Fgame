package listener

import (
	"fgame/fgame/core/event"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家进入跨服场景
func playerEnterCrossScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	crossType := pl.GetCrossType()

	activityType, ok := crossType.CrossTypeToActivityType()
	if !ok {
		return
	}
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAttandActivityPlay, int32(activityType), 1)
	return
}

func init() {
	gameevent.AddEventListener(crosseventtypes.EventTypePlayerCrossEnter, event.EventListenerFunc(playerEnterCrossScene))
}
