package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/arena"
	"fgame/fgame/cross/arena/pbutil"
	arenascene "fgame/fgame/cross/arena/scene"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家重生
func playerReborn(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArena {
		return
	}

	t := arena.GetArenaService().GetArenaTeamByPlayerId(pl.GetId())
	if t == nil {
		return
	}
	mem, _ := t.GetTeam().GetMember(pl.GetId())
	if mem == nil {
		return
	}
	if mem.GetStatus() != arenascene.MemberStatusOnline {
		return
	}
	scArenaPlayerDataRebornChanged := pbutil.BuildSCArenaPlayerDataRebornChanged(pl)
	s.BroadcastMsg(scArenaPlayerDataRebornChanged)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(playerReborn))
}
