package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	playeralliance "fgame/fgame/game/alliance/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家仙盟职位变更
func alliancePositionChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	m := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceId := m.GetAllianceId()
	allianceName := m.GetAllianceName()
	mengZhuId := m.GetMengZhuId()
	memPos := m.GetPlayerAlliancePos()

	pl.SyncAlliance(allianceId, allianceName, mengZhuId, memPos)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerAlliancePositionChanged, event.EventListenerFunc(alliancePositionChanged))
}
