package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/ring/pbutil"
	playerring "fgame/fgame/game/ring/player"
	ringtypes "fgame/fgame/game/ring/types"
)

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)

	ringObjMap := ringManager.GetPlayerRingObjectMap()
	scRingInfoGet := pbutil.BuildSCRingInfoGet(ringObjMap)
	pl.SendMsg(scRingInfoGet)

	baoKuObj := ringManager.GetPlayerBaoKuObject(ringtypes.BaoKuTypeRing)
	scRingBaoKuInfo := pbutil.BuildSCRingBaoKuInfo(baoKuObj)
	pl.SendMsg(scRingBaoKuInfo)
	return
}
