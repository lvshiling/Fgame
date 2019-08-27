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

//玩家附加系统神铸等级日志
func playerAdditionSysShenZhuLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*additionsyseventtypes.PlayerAdditionSysShenZhuLevLogEventData)
	if !ok {
		return
	}

	additionsysManager := pl.GetPlayerDataManager(playertypes.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	slotInfo := additionsysManager.GetAdditionSysByArg(eventData.GetSysType(), eventData.GetPosition())

	logAdditionSysShenZhu := &logmodel.PlayerAdditionSysShenZhu{}
	logAdditionSysShenZhu.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logAdditionSysShenZhu.BeforeLev = eventData.GetBeforeLev()
	logAdditionSysShenZhu.CurLev = slotInfo.ShenZhuLev
	logAdditionSysShenZhu.Reason = int32(eventData.GetReason())
	logAdditionSysShenZhu.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logAdditionSysShenZhu)
	return
}

func init() {
	gameevent.AddEventListener(additionsyseventtypes.EventTypeAdditionSysShenZhuLog, event.EventListenerFunc(playerAdditionSysShenZhuLog))
}
