package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/bodyshield/pbutil"
	playerbshield "fgame/fgame/game/bodyshield/player"
	gameevent "fgame/fgame/game/event"
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
	bshieldManager := pl.GetPlayerDataManager(playertypes.PlayerBShieldDataManagerType).(*playerbshield.PlayerBodyShieldDataManager)
	bShieldInfo := bshieldManager.GetBodyShiedInfo()
	scBodyShieldGet := pbutil.BuildSCBodyShieldGet(bShieldInfo)
	pl.SendMsg(scBodyShieldGet)

	scShieldGet := pbutil.BuildSCShieldGet(bShieldInfo)
	pl.SendMsg(scShieldGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
