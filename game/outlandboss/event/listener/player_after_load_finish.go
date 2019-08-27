package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/outlandboss/pbutil"
	playeroutlandboss "fgame/fgame/game/outlandboss/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerOutlandBossDataManagerType).(*playeroutlandboss.PlayerOutlandBossDataManager)
	manager.RefreshZhuoQi()

	curZhuoQi := manager.GetCurZhuoQiNum()
	scMsg := pbutil.BuildSCOutlandBossZhuoqiInfo(curZhuoQi)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
