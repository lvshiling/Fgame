package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	playeralliance "fgame/fgame/game/alliance/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家仓库积分变化
func playerDepotPointChanged(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*allianceeventtypes.PlayerAllianceDepotPointLogEventData)
	if !ok {
		return
	}

	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)

	pointChanged := &logmodel.PlayerAllianceDepotPointChanged{}
	pointChanged.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	pointChanged.ChangedPoint = eventData.GetChangedPoint()
	pointChanged.CurPoint = allianceManager.GetDepotPoint()
	pointChanged.BeforePoint = eventData.GetBeforePoint()
	pointChanged.Reason = int32(eventData.GetReason())
	pointChanged.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(pointChanged)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerAllianceDepotPointChangedLog, event.EventListenerFunc(playerDepotPointChanged))
}
