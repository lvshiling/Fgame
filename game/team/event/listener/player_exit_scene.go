package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//玩家退出场景
func playerExitScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().IsWorld() {
		return
	}
	if pl.GetTeamId() == 0 {
		return
	}
	teamData := team.GetTeamService().GetTeamByPlayerId(pl.GetId())
	if teamData == nil {
		return
	}
	scTeamMatchCondtionPrepareBroadcast := pbutil.BuildSCTeamMatchCondtionPrepareBroadcast(pl.GetId())
	teamlogic.BroadcastMsg(teamData, scTeamMatchCondtionPrepareBroadcast)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(playerExitScene))
}
