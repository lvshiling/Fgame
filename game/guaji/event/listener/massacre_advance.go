package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	massacreeventtypes "fgame/fgame/game/massacre/event/types"
	playermassacre "fgame/fgame/game/massacre/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//坐骑进阶
func massacreAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	massacreDataManager := pl.GetPlayerDataManager(playertypes.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	advanceId := int32(massacreDataManager.GetMassacreInfo().AdvanceId)
	scGuaJiAdvanceUpdateList := pbutil.BuildSCGuaJiAdvanceUpdateList(guajitypes.GuaJiAdvanceTypeMassacre, advanceId)
	pl.SendMsg(scGuaJiAdvanceUpdateList)
	return
}

func init() {
	gameevent.AddEventListener(massacreeventtypes.EventTypeMassacreAdvanced, event.EventListenerFunc(massacreAdvanced))
}
