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

//玩家点星系统升级日志
func playerDianXingAdvancedLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*dianxingeventtypes.PlayerDianXingAdvancedLogEventData)
	if !ok {
		return
	}

	dianxingManager := pl.GetPlayerDataManager(playertypes.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	dianxingInfo := dianxingManager.GetDianXingObject()

	logDianXing := &logmodel.PlayerDianXing{}
	logDianXing.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logDianXing.BeforeXingPu = eventData.GetBeforeXingPu()
	logDianXing.CurXingPu = dianxingInfo.CurrType
	logDianXing.BeforeLev = eventData.GetBeforeLev()
	logDianXing.CurLev = dianxingInfo.CurrLevel
	logDianXing.Reason = int32(eventData.GetReason())
	logDianXing.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logDianXing)
	return
}

func init() {
	gameevent.AddEventListener(dianxingeventtypes.EventTypeDianXingAdvancedLog, event.EventListenerFunc(playerDianXingAdvancedLog))
}
