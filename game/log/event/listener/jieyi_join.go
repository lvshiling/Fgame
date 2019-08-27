package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	logmodel "fgame/fgame/logserver/model"
)

//加入结义日志
func jieyiJoinJieYiLog(target event.EventTarget, data event.EventData) (err error) {

	jieYi, ok := target.(*jieyi.JieYi)
	if !ok {
		return
	}
	memberId, ok := data.(int64)
	if !ok {
		return
	}

	jieyiMerge := &logmodel.JieYiJoin{}
	jieyiMerge.JieYiLogMsg = gamelog.JieYiLogMsgFromPlayer(jieYi)
	jieyiMerge.PlayerId = memberId
	log.GetLogService().SendLog(jieyiMerge)
	return
}

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeJionJieYiLog, event.EventListenerFunc(jieyiJoinJieYiLog))
}
