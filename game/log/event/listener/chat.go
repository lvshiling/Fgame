package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/chat/chat"
	chateventtypes "fgame/fgame/game/chat/event/types"
	chattemplate "fgame/fgame/game/chat/template"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

//完成一次世界频道聊天
func playerChat(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*chat.ChatEventData)
	msgType := eventData.GetMsgType()
	if msgType == chattypes.MsgTypeText {
		content := string(eventData.GetContent())
		if !chattemplate.GetChatConstantService().IfRedirect(content) {
			return
		}
	}

	chatContent := &logmodel.ChatContent{}
	chatContent.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	chatContent.Channel = int32(eventData.GetChannel())
	chatContent.Content = eventData.GetContent()
	chatContent.MsgType = int32(eventData.GetMsgType())
	chatContent.RecvId = eventData.GetRecvId()
	recvName := eventData.GetRecvName()
	chatContent.RecvName = recvName
	if eventData.GetMsgType() == chattypes.MsgTypeText {
		chatContent.Text = string(eventData.GetContent())
	}
	log.GetLogService().SendChatLog(chatContent)
	return
}

func init() {
	gameevent.AddEventListener(chateventtypes.EventTypeChat, event.EventListenerFunc(playerChat))
}
