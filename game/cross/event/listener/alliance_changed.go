package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//仙盟变化
func allianceChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerAllianceChanged := pbutil.BuildPlayerAllianceChanged(pl.GetAllianceId(), pl.GetAllianceName(), pl.GetMengZhuId(), int32(pl.GetMemPos()))
	pl.SendCrossMsg(playerAllianceChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerAllianceChanged, event.EventListenerFunc(allianceChanged))
}
