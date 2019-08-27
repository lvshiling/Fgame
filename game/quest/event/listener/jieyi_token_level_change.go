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
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeTokenLevelChange, event.EventListenerFunc(playerJieYiTokenLevelChange))
}

func playerJieYiTokenLevelChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	level := data.(int32)

	questlogic.SetQuestData(pl, questtypes.QuestSubTypeXiongDiToken, 0, level)
	return
}
