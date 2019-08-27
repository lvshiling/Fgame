package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//从血池补血
func teShuSkillReset(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	if !pl.IsCross() {
		return
	}
	siPlayerTeShuSkillReset := pbutil.BuildSIPlayerTeShuSkillReset(pl)
	pl.SendCrossMsg(siPlayerTeShuSkillReset)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattleObjectTeshuSkillReset, event.EventListenerFunc(teShuSkillReset))
}
