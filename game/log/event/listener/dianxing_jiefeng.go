package listener

import (
	"fgame/fgame/core/event"
	dianxingeventtypes "fgame/fgame/game/dianxing/event/types"
	playerdianxing "fgame/fgame/game/dianxing/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家点星解封升级日志
func playerDianXingJieFengAdvancedLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*dianxingeventtypes.PlayerDianXingJieFengAdvancedLogEventData)
	if !ok {
		return
	}

	dianxingManager := pl.GetPlayerDataManager(playertypes.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	dianxingInfo := dianxingManager.GetDianXingObject()

	logDianXingJieFeng := &logmodel.PlayerDianXingJieFeng{}
	logDianXingJieFeng.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logDianXingJieFeng.BeforeLev = eventData.GetBeforeLev()
	logDianXingJieFeng.CurLev = dianxingInfo.JieFengLev
	logDianXingJieFeng.Reason = int32(eventData.GetReason())
	logDianXingJieFeng.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logDianXingJieFeng)
	return
}

func init() {
	gameevent.AddEventListener(dianxingeventtypes.EventTypeDianXingJieFengAdvancedLog, event.EventListenerFunc(playerDianXingJieFengAdvancedLog))
}
