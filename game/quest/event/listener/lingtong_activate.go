package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//加载完成后
func playerLingTongActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeLingTongActivateNum, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeLingTongActivate, event.EventListenerFunc(playerLingTongActivate))
}
