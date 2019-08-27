package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancelogic "fgame/fgame/game/alliance/logic"
	playeralliance "fgame/fgame/game/alliance/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
)

//TODO:zrc 修改为处理器
//仙盟改变
func allianceExit(target event.EventTarget, data event.EventData) (err error) {
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
	err = alliancelogic.ReloadAllianceSkill(pl)
	if err != nil {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeAllianceBoss && s.MapTemplate().GetMapType() != scenetypes.SceneTypeAllianceShengTan {
		return
	}
	pl.BackLastScene()

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerAllianceExit, event.EventListenerFunc(allianceExit))
}
