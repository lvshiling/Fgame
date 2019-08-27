package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//坐骑改变
func mountChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	m := pl.GetPlayerDataManager(playertypes.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountId := m.GetMountId()
	advanceId := m.GetMountAdvancedId()
	pl.SetMountId(mountId, advanceId)
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountChanged, event.EventListenerFunc(mountChanged))
}
