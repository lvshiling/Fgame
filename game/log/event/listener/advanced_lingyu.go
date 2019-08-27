package listener

import (
	"fgame/fgame/core/event"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	playerlingyu "fgame/fgame/game/lingyu/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家领域进阶日志
func playerLingyuAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*lingyueventtypes.PlayerLingyuAdvancedLogEventData)
	if !ok {
		return
	}

	lingyuManager := pl.GetPlayerDataManager(playertypes.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuInfo := lingyuManager.GetLingyuInfo()

	logLingyuAdvanced := &logmodel.PlayerLingyu{}
	logLingyuAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logLingyuAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logLingyuAdvanced.CurAdvancedNum = int32(lingyuInfo.AdvanceId)
	logLingyuAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logLingyuAdvanced.Reason = int32(eventData.GetReason())
	logLingyuAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logLingyuAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(lingyueventtypes.EventTypeLingyuAdvancedLog, event.EventListenerFunc(playerLingyuAdvancedLog))
}
