package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
	playerxianti "fgame/fgame/game/xianti/player"
)

//坐骑进阶
func xiantiAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	xiantiDataManager := pl.GetPlayerDataManager(playertypes.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	advanceId := xiantiDataManager.GetXianTiAdvancedId()
	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guajitypes.GuaJiAdvanceTypeXianti, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiAdvanced, event.EventListenerFunc(xiantiAdvanced))
}
