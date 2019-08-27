package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	"fgame/fgame/cross/arena/pbutil"
	arenascene "fgame/fgame/cross/arena/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家退出
func teamMemberExit(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	t := data.(*arenascene.ArenaTeam)
	s := pl.GetScene()
	if s == nil {
		return
	}
	switch s.MapTemplate().GetMapType() {
	case scenetypes.SceneTypeArena:
		{

			scArenaPlayerDataExitChanged := pbutil.BuildSCArenaPlayerDataExitChanged(pl)
			s.BroadcastMsg(scArenaPlayerDataExitChanged)
			if pl.IsRobot() {
				return
			}
			if t.GetState() == arenascene.ArenaTeamStateGameEnd {
				return
			}
			isMsg := pbutil.BuildISArenaGiveUp()
			pl.SendMsg(isMsg)
		}
		break
		// case scenetypes.SceneTypeArenaShengShou:
		// 	{

		// 		scArenaPlayerDataExitChanged := pbutil.BuildSCArenaPlayerDataExitChanged(pl)
		// 		arenalogic.BroadcastArenaTeam(t, scArenaPlayerDataExitChanged)
		// 	}
		// 	break
	}

	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaTeamMemberExit, event.EventListenerFunc(teamMemberExit))
}
