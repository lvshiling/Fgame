package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	ylppbutil "fgame/fgame/game/yinglingpu/pbutil"
	ylpplayer "fgame/fgame/game/yinglingpu/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	ylpManager := pl.GetPlayerDataManager(playertypes.PlayerYingLingPuManagerType).(*ylpplayer.PlayerYingLingPuManager)
	ylpList := ylpManager.GetAllYingLingPu()
	ylpSpList := ylpManager.GetAllYingLingPuSuiPian()
	sendInfo := ylppbutil.BuildYlpInfo(ylpList, ylpSpList)
	pl.SendMsg(sendInfo)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
