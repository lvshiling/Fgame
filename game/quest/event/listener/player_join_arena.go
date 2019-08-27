package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/game/arena/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家参加3v3活动活动
func playerJoinArena(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	if err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubType3V3, 0, 1); err != nil {
		return
	}
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaJoin, event.EventListenerFunc(playerJoinArena))
}
