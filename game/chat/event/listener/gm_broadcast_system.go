package listener

import (
	"fgame/fgame/core/event"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticeeventtypes "fgame/fgame/game/notice/event/types"
	"fgame/fgame/game/notice/notice"
)

//GM系统公告
func gmBroadcastSystem(target event.EventTarget, data event.EventData) (err error) {
	noticeData, ok := target.(*notice.GmNoticeData)
	if !ok {
		return
	}

	content := noticeData.GetContent()
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))

	return
}

func init() {
	gameevent.AddEventListener(noticeeventtypes.EventTypeBroadcastSystem, event.EventListenerFunc(gmBroadcastSystem))
}
