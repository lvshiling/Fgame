package listener

import (
	"fgame/fgame/core/event"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	hongbaoeventtypes "fgame/fgame/game/hongbao/event/types"
	"fgame/fgame/game/player"
	"fmt"
)

//红包发送聊天频道
func hongBapSendChat(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*hongbaoeventtypes.HongBaoSendEventData)
	hongBaoId := eventData.GetHongBaoId()
	hongBaoType := eventData.GetHongBaoType()
	hongBaoChannelType := eventData.GetHongBaoChannelType()
	args := eventData.GetCliArgs()
	if hongBaoId == 0 {
		return
	}
	if !hongBaoType.Valid() {
		return
	}
	if !hongBaoChannelType.Valid() {
		return
	}

	content := fmt.Sprintf("%d,%d", hongBaoId, int32(hongBaoType))
	switch hongBaoChannelType {
	case chattypes.ChannelTypeWorld:
		chatlogic.BroadcastHongBao(pl.GetId(), pl.GetName(), []byte(content), args)
		break
	case chattypes.ChannelTypeBangPai:
		allianceId := pl.GetAllianceId()
		if allianceId == 0 {
			return
		}
		chatlogic.BroadcastAlliance(allianceId, pl.GetId(), pl.GetName(), chattypes.MsgTypeHongBao, []byte(content), args)
		break
	}
	return
}

func init() {
	gameevent.AddEventListener(hongbaoeventtypes.EventTypeHongBaoSend, event.EventListenerFunc(hongBapSendChat))
}
