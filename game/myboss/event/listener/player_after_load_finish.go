package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/myboss/pbutil"
	mybossplayer "fgame/fgame/game/myboss/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	//获取信息List
	mybossManager := pl.GetPlayerDataManager(types.PlayerMyBossDataManagerType).(*mybossplayer.PlayerMyBossDataManager)
	timesAll := mybossManager.GetAttendTimesAll()
	scMsg := pbutil.BuildSCMyBossInfoNotice(timesAll)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
