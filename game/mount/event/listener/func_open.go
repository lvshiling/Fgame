package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//坐骑改变
func mountFuncOpen(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	funcType := data.(funcopentypes.FuncOpenType)

	if funcType != funcopentypes.FuncOpenTypeMount {
		return
	}
	m := pl.GetPlayerDataManager(playertypes.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountId := m.GetMountId()
	advanceId := m.GetMountAdvancedId()
	pl.SetMountId(mountId, advanceId)
	return
}

func init() {
	gameevent.AddEventListener(funcopeneventtypes.EventTypeFuncOpen, event.EventListenerFunc(mountFuncOpen))
}
