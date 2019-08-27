package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/game/arena/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家3V3胜利次数
func playerArenaWinChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	win, ok := data.(bool)
	if !ok {
		return
	}
	if !win {
		return
	}

	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubType3V3WinNum, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaWinChanged, event.EventListenerFunc(playerArenaWinChanged))
}
