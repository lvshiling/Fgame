package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	feishengeventtypes "fgame/fgame/game/feisheng/event/types"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家飞升日志
func playerFeiShengLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*feishengeventtypes.PlayerFeiShengLogEventData)
	if !ok {
		return
	}

	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiInfo := feiManager.GetFeiShengInfo()
	beforeLevel := eventData.GetBeforeLevel()
	curLevel := feiInfo.GetFeiLevel()
	beforeGongDe := eventData.GetBeforeGongDe()
	curGongDe := feiInfo.GetGongDeNum()
	uplevel := curLevel - beforeLevel
	costGongDe := beforeGongDe - curGongDe

	logFeiSheng := &logmodel.PlayerFeiSheng{}
	logFeiSheng.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logFeiSheng.CurFeiShengLevel = curLevel
	logFeiSheng.BeforeFeiShengLevel = beforeLevel
	logFeiSheng.Uplevel = uplevel
	logFeiSheng.BeforeGongDe = beforeGongDe
	logFeiSheng.CurGongDe = curGongDe
	logFeiSheng.CostGongDe = costGongDe
	logFeiSheng.Reason = int32(eventData.GetReason())
	logFeiSheng.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logFeiSheng)
	return
}

func init() {
	gameevent.AddEventListener(feishengeventtypes.EventTypePlayerFeiShengLog, event.EventListenerFunc(playerFeiShengLog))
}
