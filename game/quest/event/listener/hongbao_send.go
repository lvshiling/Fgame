package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	hongbaoeventtypes "fgame/fgame/game/hongbao/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//红包发送聊天频道
func hongBapSendChat(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeSendHongBao, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(hongbaoeventtypes.EventTypeHongBaoSend, event.EventListenerFunc(hongBapSendChat))
}
