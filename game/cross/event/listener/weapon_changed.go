package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//战翼改变
func weaponChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerWeaponChanged := pbutil.BuildPlayerWeaponChanged(pl.GetWeaponId(), pl.GetWeaponState())
	pl.SendCrossMsg(playerWeaponChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowWeaponChanged, event.EventListenerFunc(weaponChanged))
}
