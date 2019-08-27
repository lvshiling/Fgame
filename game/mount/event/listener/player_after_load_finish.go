package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/mount/pbutil"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	mountManager := p.GetPlayerDataManager(playertypes.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountInfo := mountManager.GetMountInfo()
	mountOtherMap := mountManager.GetMountOtherMap()
	scMountGet := pbutil.BuildSCMountGet(mountInfo, mountOtherMap)
	p.SendMsg(scMountGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
