package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/battle/battle"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//灵童变化
func playerLingTongChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	eventData := data.(*battle.LingTongChangeEventData)

	oldLingTong := eventData.GetOldLingTong()
	if oldLingTong != nil {
		lingTongRemove := pbutil.BuildLingTongRemove()
		pl.SendCrossMsg(lingTongRemove)
	}
	newLingTong := eventData.GetNewLingTong()
	if newLingTong != nil {
		lingTongDataInit := pbutil.BuildLingTongDataInit(newLingTong)
		pl.SendCrossMsg(lingTongDataInit)
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerLingTongChanged, event.EventListenerFunc(playerLingTongChanged))
}
