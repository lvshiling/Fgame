package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/emperor/emperor"
	"fgame/fgame/game/emperor/pbutil"
	playeremperor "fgame/fgame/game/emperor/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerEmperorDataManagerType).(*playeremperor.PlayerEmperorDataManager)
	emperorObj := emperor.GetEmperorService().GetEmperorInfo()
	worshipNum := manager.GetWorshipNum()
	scEmperorGet := pbuitl.BuildSCEmperorGet(emperorObj, worshipNum)
	pl.SendMsg(scEmperorGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
