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

//玩家附加系统觉醒日志
func playerAdditionSysAwakeLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*additionsyseventtypes.PlayerAdditionSysAwakeEventData)
	if !ok {
		return
	}

	additionsysManager := pl.GetPlayerDataManager(playertypes.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	AwakeInfo := additionsysManager.GetAdditionSysAwakeInfoByType(eventData.GetSysType())

	logAdditionSysAwake := &logmodel.PlayerAdditionSysAwake{}
	logAdditionSysAwake.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logAdditionSysAwake.BeforeIsAwake = eventData.GetBeforeIsAwake()
	logAdditionSysAwake.CurIsAwake = AwakeInfo.IsAwake
	logAdditionSysAwake.Reason = int32(eventData.GetReason())
	logAdditionSysAwake.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logAdditionSysAwake)
	return
}

func init() {
	gameevent.AddEventListener(additionsyseventtypes.EventTypeAdditionSysAwakeLog, event.EventListenerFunc(playerAdditionSysAwakeLog))
}
