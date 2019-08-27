package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	xianTiManager := p.GetPlayerDataManager(playertypes.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xianTiInfo := xianTiManager.GetXianTiInfo()
	xianTiOtherMap := xianTiManager.GetXianTiOtherMap()
	scXianTiGet := pbutil.BuildSCXianTiGet(xianTiInfo, xianTiOtherMap)
	p.SendMsg(scXianTiGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
