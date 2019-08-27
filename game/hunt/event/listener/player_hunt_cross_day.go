package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	hunteventtypes "fgame/fgame/game/hunt/event/types"
	"fgame/fgame/game/hunt/pbutil"
	playerhunt "fgame/fgame/game/hunt/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//寻宝跨天刷新
func playerHuntCrossDa(target event.EventTarget, data event.EventData) (err error) {
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
	gameevent.AddEventListener(hunteventtypes.EventTypeHuntCrossDay, event.EventListenerFunc(playerHuntCrossDa))
}
