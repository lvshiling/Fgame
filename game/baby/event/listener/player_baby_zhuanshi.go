package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/baby/baby"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

// 玩家宝宝转世
func playerBabyZhuanShi(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	babyId, ok := data.(int64)
	if !ok {
		return
	}

	baby.GetBabyService().BabyZhuanShi(pl, babyId)
	return
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyZhuanShi, event.EventListenerFunc(playerBabyZhuanShi))
}
