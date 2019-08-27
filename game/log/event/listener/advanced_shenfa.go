package listener

import (
	"fgame/fgame/core/event"
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	playershenfa "fgame/fgame/game/shenfa/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家身法进阶日志
func playerShenfaAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*shenfaeventtypes.PlayerShenfaAdvancedLogEventData)
	if !ok {
		return
	}

	shenfaManager := pl.GetPlayerDataManager(playertypes.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenfaInfo := shenfaManager.GetShenfaInfo()

	logShenfaAdvanced := &logmodel.PlayerShenfa{}
	logShenfaAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logShenfaAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logShenfaAdvanced.CurAdvancedNum = int32(shenfaInfo.AdvanceId)
	logShenfaAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logShenfaAdvanced.Reason = int32(eventData.GetReason())
	logShenfaAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logShenfaAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaAdvancedLog, event.EventListenerFunc(playerShenfaAdvancedLog))
}
