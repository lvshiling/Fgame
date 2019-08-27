package listener

import (
	"fgame/fgame/core/event"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家护体盾进阶日志
func playerBodyShieldAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*bodyshieldeventtypes.PlayerBodyShieldAdvancedLogEventData)
	if !ok {
		return
	}

	bodyshieldManager := pl.GetPlayerDataManager(playertypes.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	bodyshieldInfo := bodyshieldManager.GetBodyShiedInfo()

	logBodyShieldAdvanced := &logmodel.PlayerBodyShield{}
	logBodyShieldAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logBodyShieldAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logBodyShieldAdvanced.CurAdvancedNum = int32(bodyshieldInfo.AdvanceId)
	logBodyShieldAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logBodyShieldAdvanced.Reason = int32(eventData.GetReason())
	logBodyShieldAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logBodyShieldAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(bodyshieldeventtypes.EventTypeBodyShieldAdvancedLog, event.EventListenerFunc(playerBodyShieldAdvancedLog))
}
