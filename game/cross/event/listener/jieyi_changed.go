package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//结义变化
func jieYiChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerJieYiChanged := pbutil.BuildPlayerJieYiChanged(pl.GetJieYiId(), pl.GetJieYiName(), pl.GetJieYiRank())
	pl.SendCrossMsg(playerJieYiChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerJieYiChanged, event.EventListenerFunc(jieYiChanged))
}
