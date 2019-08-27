package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	playerwing "fgame/fgame/game/wing/player"
)

//坐骑进阶
func featherAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	wingDataManager := pl.GetPlayerDataManager(playertypes.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	advanceId := wingDataManager.GetWingInfo().FeatherId
	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guajitypes.GuaJiAdvanceTypeFeather, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeFeatherAdvanced, event.EventListenerFunc(featherAdvanced))
}
