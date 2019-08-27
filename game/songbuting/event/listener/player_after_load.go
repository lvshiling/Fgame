package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/songbuting/pbutil"
	playersongbuting "fgame/fgame/game/songbuting/player"
)

func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	manager := pl.GetPlayerDataManager(playertypes.PlayerSongBuTingDataManagerType).(*playersongbuting.PlayerSongBuTingDataManager)
	songBuTingObj := manager.GetSongBuTingObj()
	scSongBuTingChanged := pbutil.BuildSCSongBuTingChanged(songBuTingObj)
	pl.SendMsg(scSongBuTingChanged)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
