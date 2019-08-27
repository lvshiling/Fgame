package listener

import (
	"fgame/fgame/core/event"
	chesseventtypes "fgame/fgame/game/chess/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家苍龙棋局抽奖
func playerChessRew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeDragonChess, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(chesseventtypes.EventTypeAttendChess, event.EventListenerFunc(playerChessRew))
}
