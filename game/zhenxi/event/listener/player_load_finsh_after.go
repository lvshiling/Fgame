package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	zhenxipbutil "fgame/fgame/game/zhenxi/pbutil"
	playerzhenxi "fgame/fgame/game/zhenxi/player"
)

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}

// 玩家加载
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	zhenXinManager := pl.GetPlayerDataManager(playertypes.PlayerZhenXiDataManagerType).(*playerzhenxi.PlayerZhenXiDataManager)
	obj := zhenXinManager.GetPlayerZhenXiObject()
	scPlayerZhenXiBossInfo := zhenxipbutil.BuildSCPlayerZhenXiBossInfo(obj.GetEnterTimes())
	pl.SendMsg(scPlayerZhenXiBossInfo)

	return
}
