package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//坐骑进阶
func mountAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	mountManager := pl.GetPlayerDataManager(playertypes.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	advanceId := mountManager.GetMountAdvancedId()
	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guajitypes.GuaJiAdvanceTypeMount, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountAdvanced, event.EventListenerFunc(mountAdvanced))
}
