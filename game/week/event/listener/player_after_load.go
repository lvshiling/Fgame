package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/week/pbutil"
	playerweek "fgame/fgame/game/week/player"
)

// 玩家加载后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	//下发周卡信息
	weekManager := pl.GetPlayerDataManager(playertypes.PlayerWeekDataManagerType).(*playerweek.PlayerWeekManager)

	weekInfoMap := weekManager.GetWeekInfoMap()
	scMsg := pbutil.BuildSCWeekInfo(weekInfoMap)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
