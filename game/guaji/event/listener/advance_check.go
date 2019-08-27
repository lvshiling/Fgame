package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	guajieventtypes "fgame/fgame/game/guaji/event/types"
	"fgame/fgame/game/guaji/guaji"
	playerguaji "fgame/fgame/game/guaji/player"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//检查挂机提升
func guaJiAdvanceCheck(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	guaJiManager := pl.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)
	autoBuyValue := guaJiManager.GetGlobalValue(guajitypes.GuaJiGlobalTypeAdvanceAutoBuy)
	autoBuy := true
	if autoBuyValue == 0 {
		autoBuy = false
	}
	//TODO 排序
	for t, h := range guaji.GetGuaJiAdvanceCheckHandlerMap() {
		advanceId := guaJiManager.GetAdvanceSettingValue(t)
		h.AdvanceCheck(pl, t, advanceId, autoBuy)
	}
	return
}

func init() {
	gameevent.AddEventListener(guajieventtypes.GuaJiEventTypeGuaJiAdvanceCheck, event.EventListenerFunc(guaJiAdvanceCheck))
}
