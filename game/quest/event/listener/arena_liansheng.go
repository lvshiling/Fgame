package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/game/arena/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家3V3层数
func playerArenaLianSheng(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	winCount := data.(int32)

	err = questlogic.SetQuestDataSurpass(pl, questtypes.QuestSubType3V3LianSheng, 0, winCount)
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaLianSheng, event.EventListenerFunc(playerArenaLianSheng))
}
