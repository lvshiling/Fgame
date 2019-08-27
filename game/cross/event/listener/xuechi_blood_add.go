package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/xuechi/pbutil"
)

//从血池补血
func xueChiAdd(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	addBlood := data.(int64)
	if !pl.IsCross() {
		return
	}
	siXueChiAdd := pbutil.BuildSIXueChiAdd(addBlood)
	pl.SendCrossMsg(siXueChiAdd)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerXueChiBloodAdd, event.EventListenerFunc(xueChiAdd))
}
