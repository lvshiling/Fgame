package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	wingManager := pl.GetPlayerDataManager(playertypes.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	wingTrialInfo := wingManager.GetWingTrialInfo()
	wingOtherMap := wingManager.GetWingOtherMap()
	scWingGet := pbutil.BuildSCWingGet(wingInfo, wingTrialInfo, wingOtherMap)
	pl.SendMsg(scWingGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
