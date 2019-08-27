package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/fashion/pbutil"
	playerfashion "fgame/fgame/game/fashion/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	fashionManager := pl.GetPlayerDataManager(playertypes.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	fashionWear := fashionManager.GetFashionId()
	fashionMap := fashionManager.GetFashionMap()
	trialMap := fashionManager.GetTrialFashionMap()
	scfashionGet := pbutil.BuildSCFashionGet(fashionWear, fashionMap, trialMap)
	pl.SendMsg(scfashionGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
