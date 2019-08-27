package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	"fgame/fgame/cross/arena/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家离线
func teamMemberOffline(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	scArenaPlayerDataOfflineChanged := pbutil.BuildSCArenaPlayerDataOfflineChanged(pl)
	switch s.MapTemplate().GetMapType() {
	case scenetypes.SceneTypeArena:
		s.BroadcastMsg(scArenaPlayerDataOfflineChanged)
		break
		// case scenetypes.SceneTypeArenaShengShou:
		// 	t := data.(*arenascene.ArenaTeam)
		// 	arenalogic.BroadcastArenaTeam(t, scArenaPlayerDataOfflineChanged)
		// 	break
	}

	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaTeamMemberOffline, event.EventListenerFunc(teamMemberOffline))
}
