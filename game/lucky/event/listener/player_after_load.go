package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/lucky/pbutil"
	playerlucky "fgame/fgame/game/lucky/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

// 玩家加载后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	luckyManager := pl.GetPlayerDataManager(playertypes.PlayerLuckyDataManagerType).(*playerlucky.PlayerLuckyDataManager)

	luckyMap := luckyManager.GetAllLuckyObj()
	if len(luckyMap) > 0 {
		scMsg := pbutil.BuildSCLuckyInfoNotice(luckyMap)
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
