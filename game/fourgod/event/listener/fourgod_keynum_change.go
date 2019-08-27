package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	"fgame/fgame/game/fourgod/pbutil"
	playerfourgod "fgame/fgame/game/fourgod/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//四神遗迹玩家钥匙改变
func fourGodKeyNumChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	keyNum := manager.GetKeyNum()
	scFourGodKeyNumChange := pbuitl.BuildSCFourGodKeyNumChange(keyNum)
	pl.SendMsg(scFourGodKeyNumChange)
	pl.SetFourGodKey(keyNum)
	return
}

func init() {
	gameevent.AddEventListener(fourgodeventtypes.EventTypeFourGodKeyChange, event.EventListenerFunc(fourGodKeyNumChange))
}
