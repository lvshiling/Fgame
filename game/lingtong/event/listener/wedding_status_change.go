package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/scene/scene"
)

//结婚状态改变
func weddingStatusChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}

	//巡游阶段隐藏灵童
	if pl.GetWeddingStatus() == int32(marrytypes.MarryWedStatusSelfTypeCruise) {
		pl.HiddenLingTong(true)
	} else {
		pl.HiddenLingTong(false)
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowWeddingStatusChanged, event.EventListenerFunc(weddingStatusChanged))
}
