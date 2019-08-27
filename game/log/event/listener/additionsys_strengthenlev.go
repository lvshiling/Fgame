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

//玩家附加系统强化等级日志
func playerAdditionSysStrengthenLevLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*additionsyseventtypes.PlayerAdditionSysStrengthenLevLogEventData)
	if !ok {
		return
	}

	additionsysManager := pl.GetPlayerDataManager(playertypes.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	slotInfo := additionsysManager.GetAdditionSysByArg(eventData.GetSysType(), eventData.GetPosition())

	logAdditionSysStrengthenLev := &logmodel.PlayerAdditionSysStrengthenLev{}
	logAdditionSysStrengthenLev.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logAdditionSysStrengthenLev.BeforeLev = eventData.GetBeforeLev()
	logAdditionSysStrengthenLev.CurLev = slotInfo.Level
	logAdditionSysStrengthenLev.Reason = int32(eventData.GetReason())
	logAdditionSysStrengthenLev.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logAdditionSysStrengthenLev)
	return
}

func init() {
	gameevent.AddEventListener(additionsyseventtypes.EventTypeAdditionSysStrengthenLevLog, event.EventListenerFunc(playerAdditionSysStrengthenLevLog))
}
