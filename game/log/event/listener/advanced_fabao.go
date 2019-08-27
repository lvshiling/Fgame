package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	playerfabao "fgame/fgame/game/fabao/player"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家法宝进阶日志
func playerFaBaoAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*fabaoeventtypes.PlayerFaBaoAdvancedLogEventData)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	faBaoInfo := manager.GetFaBaoInfo()
	logFaBaoAdvanced := &logmodel.PlayerFaBao{}
	logFaBaoAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logFaBaoAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logFaBaoAdvanced.CurAdvancedNum = faBaoInfo.GetAdvancedId()
	logFaBaoAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logFaBaoAdvanced.Reason = int32(eventData.GetReason())
	logFaBaoAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logFaBaoAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(fabaoeventtypes.EventTypeFaBaoAdvancedLog, event.EventListenerFunc(playerFaBaoAdvancedLog))
}
