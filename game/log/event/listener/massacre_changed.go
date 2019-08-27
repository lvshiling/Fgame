package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	massacreeventtypes "fgame/fgame/game/massacre/event/types"
	playermassacre "fgame/fgame/game/massacre/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家戮仙刃进阶日志
func playerMassacreChangedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*massacreeventtypes.PlayerMassacreChangedLogEventData)
	if !ok {
		return
	}

	s := pl.GetScene()
	mapId := int32(0)
	if s != nil {
		mapId = s.MapId()
	}
	massacreManager := pl.GetPlayerDataManager(playertypes.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	massacreInfo := massacreManager.GetMassacreInfo()

	logMassacreChanged := &logmodel.PlayerMassacre{}
	logMassacreChanged.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logMassacreChanged.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logMassacreChanged.CurAdvancedNum = int32(massacreInfo.AdvanceId)
	logMassacreChanged.ChangedNum = eventData.GetChangedNum()
	logMassacreChanged.BefShaQiNum = eventData.GetBeforeShaQiNum()
	logMassacreChanged.CurShaQiNum = massacreInfo.ShaQiNum
	logMassacreChanged.CurMapId = mapId
	logMassacreChanged.Reason = int32(eventData.GetReason())
	logMassacreChanged.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logMassacreChanged)
	return
}

func init() {
	gameevent.AddEventListener(massacreeventtypes.EventTypeMassacreChangedLog, event.EventListenerFunc(playerMassacreChangedLog))
}
