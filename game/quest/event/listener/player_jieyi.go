package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

// 玩家结义
func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeJieYiSuccess, event.EventListenerFunc(playerJieYiSuccess))
}

func playerJieYiSuccess(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFinishJieYi, 0, 1)
	return
}
