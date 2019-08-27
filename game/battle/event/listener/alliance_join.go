package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancelogic "fgame/fgame/game/alliance/logic"
	playeralliance "fgame/fgame/game/alliance/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//仙盟改变
func allianceJoin(target event.EventTarget, data event.EventData) (err error) {
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

	alliancelogic.ReloadAllianceSkill(pl)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerAllianceJoin, event.EventListenerFunc(allianceJoin))
}
