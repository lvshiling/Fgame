package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家噬魂幡进阶日志
func playerShiHunFanAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*shihunfaneventtypes.PlayerShiHunFanAdvancedLogEventData)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shihunfanInfo := manager.GetShiHunFanInfo()

	logShiHunFanAdvanced := &logmodel.PlayerShiHunFan{}
	logShiHunFanAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logShiHunFanAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logShiHunFanAdvanced.CurAdvancedNum = int32(shihunfanInfo.AdvanceId)
	logShiHunFanAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logShiHunFanAdvanced.Reason = int32(eventData.GetReason())
	logShiHunFanAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logShiHunFanAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(shihunfaneventtypes.EventTypeShiHunFanAdvancedLog, event.EventListenerFunc(playerShiHunFanAdvancedLog))
}
