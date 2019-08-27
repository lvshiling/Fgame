package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家灵童养成类进阶日志
func playerLingTongDevAdvancedLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*lingtongdeveventtypes.PlayerLingTongDevAdvancedLogEventData)
	if !ok {
		return
	}
	classType := eventData.GetClassType()

	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongInfo := manager.GetLingTongDevInfo(classType)
	if lingTongInfo == nil {
		return
	}

	logLingTongDevAdvanced := &logmodel.PlayerLingTongDev{}
	logLingTongDevAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logLingTongDevAdvanced.ClassType = int32(classType)
	logLingTongDevAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logLingTongDevAdvanced.CurAdvancedNum = lingTongInfo.GetAdvancedId()
	logLingTongDevAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logLingTongDevAdvanced.Reason = int32(eventData.GetReason())
	logLingTongDevAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logLingTongDevAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevAdvancedLog, event.EventListenerFunc(playerLingTongDevAdvancedLog))
}
