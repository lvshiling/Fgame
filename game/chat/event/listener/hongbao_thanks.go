package listener

import (
	lang "fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	hongbaoeventtypes "fgame/fgame/game/hongbao/event/types"
	"fgame/fgame/game/player"
)

//红包感谢发言
func hongBapThanksChat(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	channelType := data.(chattypes.ChannelType)
	content := lang.GetLangService().ReadLang(lang.HongBaoSnatchThanksBoss)
	chatlogic.BroadcastHongBaoThanks(pl, channelType, []byte(content), "")
	return
}

func init() {
	gameevent.AddEventListener(hongbaoeventtypes.EventTypeHongBaoThanks, event.EventListenerFunc(hongBapThanksChat))
}
