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
func playerAdditionSysShengJiLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*additionsyseventtypes.PlayerAdditionSysShengJiLogEventData)
	if !ok {
		return
	}

	additionsysManager := pl.GetPlayerDataManager(playertypes.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	LevelInfo := additionsysManager.GetAdditionSysLevelInfoByType(eventData.GetSysType())

	logAdditionSysShengJi := &logmodel.PlayerAdditionSysShengJi{}
	logAdditionSysShengJi.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logAdditionSysShengJi.BeforeLev = eventData.GetBeforeLev()
	logAdditionSysShengJi.CurLev = LevelInfo.Level
	logAdditionSysShengJi.BeforeUpNum = eventData.GetBeforeUpNum()
	logAdditionSysShengJi.CurUpNum = LevelInfo.UpNum
	logAdditionSysShengJi.BeforeUpPro = eventData.GetBeforeUpPro()
	logAdditionSysShengJi.CurUpPro = LevelInfo.UpPro
	logAdditionSysShengJi.Reason = int32(eventData.GetReason())
	logAdditionSysShengJi.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logAdditionSysShengJi)
	return
}

func init() {
	gameevent.AddEventListener(additionsyseventtypes.EventTypeAdditionSysShengJiLog, event.EventListenerFunc(playerAdditionSysShengJiLog))
}
