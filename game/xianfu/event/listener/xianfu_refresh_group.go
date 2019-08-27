package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
	xianfuplayer "fgame/fgame/game/xianfu/player"
)

//更新副本波数
func playerXianFuRefreshGroup(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	group := data.(int32)
	xianfuManager := pl.GetPlayerDataManager(types.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	xianfuManager.RefreshGroup(group)

	return
}

func init() {
	gameevent.AddEventListener(xianfueventtypes.EventTypeXianFuRefreshGroup, event.EventListenerFunc(playerXianFuRefreshGroup))
}
