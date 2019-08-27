package listener

import (
	"fgame/fgame/core/event"
	additionsyseventtypes "fgame/fgame/game/additionsys/event/types"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家附加系统通灵等级日志
func playerAdditionSysTongLingLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*additionsyseventtypes.PlayerAdditionSysTongLingLogEventData)
	if !ok {
		return
	}

	additionsysManager := pl.GetPlayerDataManager(playertypes.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	tongLingInfo := additionsysManager.GetAdditionSysTongLingInfoByType(eventData.GetSysType())

	logAdditionSysTongLing := &logmodel.PlayerAdditionSysTongLing{}
	logAdditionSysTongLing.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logAdditionSysTongLing.BeforeLev = eventData.GetBeforeLev()
	logAdditionSysTongLing.CurLev = tongLingInfo.TongLingLev
	logAdditionSysTongLing.Reason = int32(eventData.GetReason())
	logAdditionSysTongLing.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logAdditionSysTongLing)
	return
}

func init() {
	gameevent.AddEventListener(additionsyseventtypes.EventTypeAdditionSysTongLingLog, event.EventListenerFunc(playerAdditionSysTongLingLog))
}
