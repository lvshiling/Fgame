package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/baby/baby"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	playerbaby "fgame/fgame/game/baby/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

// 玩家宝宝读书升级
func playerBabyLearnUplevel(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	babyObj, ok := data.(*playerbaby.PlayerBabyObject)
	if !ok {
		return
	}

	// 同步到全局表
	baby.GetBabyService().SyncBabyLearnLevel(pl, babyObj.GetDBId(), babyObj.GetLearnLevel())
	return
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyLearnUplevel, event.EventListenerFunc(playerBabyLearnUplevel))
}
