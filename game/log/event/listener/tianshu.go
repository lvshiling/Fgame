package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	tianshueventtypes "fgame/fgame/game/tianshu/event/types"
	playertianshu "fgame/fgame/game/tianshu/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家天书日志
func playerTianShuAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*tianshueventtypes.PlayerTianShuLogEventData)
	if !ok {
		return
	}

	tianshuManager := pl.GetPlayerDataManager(playertypes.PlayerTianShuDataManagerType).(*playertianshu.PlayerTianShuDataManager)
	curLevel := tianshuManager.GetTianShuLevel(eventData.GetTianShuType())

	logTianShu := &logmodel.PlayerTianShu{}
	logTianShu.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logTianShu.BeforeLevel = eventData.GetBeforeLevel()
	logTianShu.Uplevel = eventData.GetUpLevel()
	logTianShu.CurLevel = curLevel
	logTianShu.Reason = int32(eventData.GetReason())
	logTianShu.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logTianShu)
	return
}

func init() {
	gameevent.AddEventListener(tianshueventtypes.EventTypePlayerTianShuLog, event.EventListenerFunc(playerTianShuAdvancedLog))
}
