package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/chat/chat"
	"fgame/fgame/game/chat/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)

	var worldChatList []*chat.ChatData
	systemChatList := chat.GetChatService().GetSystemChatList()
	scWorldChatListNotice := pbutil.BuildSCWorldChatListNotice(worldChatList, systemChatList)
	p.SendMsg(scWorldChatListNotice)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
