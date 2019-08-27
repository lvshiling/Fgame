package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//队伍变化
func teamChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerTeamChanged := pbutil.BuildPlayerTeamChanged(pl.GetTeamId(), pl.GetTeamName(), int32(pl.GetTeamPurpose()))
	pl.SendCrossMsg(playerTeamChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerTeamChanged, event.EventListenerFunc(teamChanged))
}
