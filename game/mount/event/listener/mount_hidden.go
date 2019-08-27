package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//坐骑隐藏或穿戴
func mountHidden(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	hiddenMount := pl.IsMountHidden()
	manager := pl.GetPlayerDataManager(playertypes.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	manager.Hidden(hiddenMount)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowMountHidden, event.EventListenerFunc(mountHidden))
}
