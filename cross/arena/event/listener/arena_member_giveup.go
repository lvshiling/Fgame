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

//玩家放弃
func arenaTeamMemberGiveUp(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	t := data.(*arenascene.ArenaTeam)
	s := pl.GetScene()
	if s == nil {
		return
	}
	switch s.MapTemplate().GetMapType() {
	case scenetypes.SceneTypeArena:
		{

			scArenaPlayerDataFailedChanged := pbutil.BuildSCArenaPlayerDataFailedChanged(pl)
			s.BroadcastMsg(scArenaPlayerDataFailedChanged)
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

		// 		scArenaPlayerDataFailedChanged := pbutil.BuildSCArenaPlayerDataFailedChanged(pl)
		// 		arenalogic.BroadcastArenaTeam(t, scArenaPlayerDataFailedChanged)
		// 	}
		// 	break
	}

	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaTeamMemberGiveUp, event.EventListenerFunc(arenaTeamMemberGiveUp))
}
