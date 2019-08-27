package listener

import (
	"fgame/fgame/core/event"
	anqieventtypes "fgame/fgame/game/anqi/event/types"
	playeranqi "fgame/fgame/game/anqi/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//坐骑进阶
func anqiAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	anqiDataManager := pl.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	advanceId := anqiDataManager.GetAnqiAdvanced()
	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guajitypes.GuaJiAdvanceTypeAnqi, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(anqieventtypes.EventTypeAnqiAdvanced, event.EventListenerFunc(anqiAdvanced))
}
