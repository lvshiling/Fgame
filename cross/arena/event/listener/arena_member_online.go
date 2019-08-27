package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	"fgame/fgame/cross/arena/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家在线
func teamMemberOnline(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	scArenaPlayerDataOnlineChanged := pbutil.BuildSCArenaPlayerDataOnlineChanged(pl)
	switch s.MapTemplate().GetMapType() {
	case scenetypes.SceneTypeArena:
		s.BroadcastMsg(scArenaPlayerDataOnlineChanged)
		break
		// case scenetypes.SceneTypeArenaShengShou:
		// 	t := data.(*arenascene.ArenaTeam)
		// 	arenalogic.BroadcastArenaTeam(t, scArenaPlayerDataOnlineChanged)
		// 	break
	}

	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaTeamMemberOnline, event.EventListenerFunc(teamMemberOnline))
}
