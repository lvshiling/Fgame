package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
)

//玩家队伍改变
func playerTeamChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	if s == nil {
		return
	}
	teamId := pl.GetTeamId()
	if teamId == 0 {
		if pl.GetPkState() == pktypes.PkStateGroup {
			defaultPkState := s.MapTemplate().GetPkState()
			pl.SwitchPkState(defaultPkState, pktypes.PkCommonCampDefault)
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerTeamChanged, event.EventListenerFunc(playerTeamChanged))
}
