package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	playerfabao "fgame/fgame/game/fabao/player"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//坐骑进阶
func fabaoAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fabaoDataManager := pl.GetPlayerDataManager(playertypes.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	advanceId := fabaoDataManager.GetFaBaoAdvancedId()
	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guajitypes.GuaJiAdvanceTypeFabao, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(fabaoeventtypes.EventTypeFaBaoAdvanced, event.EventListenerFunc(fabaoAdvanced))
}
