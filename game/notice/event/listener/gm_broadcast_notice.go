package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	noticeeventtypes "fgame/fgame/game/notice/event/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/notice/notice"
)

//GM公告
func gmBroadcastNotice(target event.EventTarget, data event.EventData) (err error) {
	noticeData, ok := target.(*notice.GmNoticeData)
	if !ok {
		return
	}

	content := noticeData.GetContent()
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(noticeeventtypes.EventTypeBroadcastNotice, event.EventListenerFunc(gmBroadcastNotice))
}
