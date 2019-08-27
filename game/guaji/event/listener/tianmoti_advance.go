package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	playertianmo "fgame/fgame/game/tianmo/player"
)

//坐骑进阶
func tianmoAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	tianmoDataManager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	advanceId := tianmoDataManager.GetTianMoAdvanced()
	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guajitypes.GuaJiAdvanceTypeTianmoti, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(tianmoeventtypes.EventTypeTianMoAdvanced, event.EventListenerFunc(tianmoAdvanced))
}
