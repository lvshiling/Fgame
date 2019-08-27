package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	playershenfa "fgame/fgame/game/shenfa/player"
)

//坐骑进阶
func shenfaAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	shenfaDataManager := pl.GetPlayerDataManager(playertypes.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	advanceId := shenfaDataManager.GetShenfaAdvanced()
	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guajitypes.GuaJiAdvanceTypeShenfa, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaAdvanced, event.EventListenerFunc(shenfaAdvanced))
}
