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

//玩家盾刺进阶日志
func playerShieldAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*bodyshieldeventtypes.PlayerShieldAdvancedLogEventData)
	if !ok {
		return
	}

	shieldManager := pl.GetPlayerDataManager(playertypes.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	shieldInfo := shieldManager.GetBodyShiedInfo()

	logShieldAdvanced := &logmodel.PlayerShield{}
	logShieldAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logShieldAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logShieldAdvanced.CurAdvancedNum = int32(shieldInfo.ShieldId)
	logShieldAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logShieldAdvanced.Reason = int32(eventData.GetReason())
	logShieldAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logShieldAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(bodyshieldeventtypes.EventTypeShieldAdvancedLog, event.EventListenerFunc(playerShieldAdvancedLog))
}
