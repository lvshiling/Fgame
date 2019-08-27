package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/hunt/pbutil"
	playerhunt "fgame/fgame/game/hunt/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	huntManager := pl.GetPlayerDataManager(playertypes.PlayerHuntDataManagerType).(*playerhunt.PlayerHuntDataManager)
	huntInfoMap := huntManager.GetAllHuntInfo()
	scMsg := pbutil.BuildSCHuntInfoNotice(huntInfoMap)
	pl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
