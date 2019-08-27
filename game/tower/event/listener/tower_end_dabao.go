package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	towerventtypes "fgame/fgame/game/tower/event/types"
	"fgame/fgame/game/tower/pbutil"
	playertower "fgame/fgame/game/tower/player"
)

//结束塔跨天
func endDaBao(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	pl.EndDaBao()

	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)
	remainTime := towerManager.GetRemainTime()
	scMsg := pbutil.BuildSCTowerTimeNotice(remainTime)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(towerventtypes.EventTypeTowerEndDaBao, event.EventListenerFunc(endDaBao))
}
