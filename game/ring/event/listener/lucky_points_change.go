package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	ringeventtypes "fgame/fgame/game/ring/event/types"
	"fgame/fgame/game/ring/pbutil"
	playerring "fgame/fgame/game/ring/player"
	ringtypes "fgame/fgame/game/ring/types"
)

func init() {
	gameevent.AddEventListener(ringeventtypes.EventTypeRingLuckyPointsChange, event.EventListenerFunc(luckyPointsChange))
}

func luckyPointsChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	typ, ok := data.(ringtypes.BaoKuType)
	if !ok {
		return
	}

	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)
	baoKuObj := ringManager.GetPlayerBaoKuObject(typ)
	scRingLuckyPointsChange := pbutil.BuildSCRingLuckyPointsChange(int32(baoKuObj.GetType()), baoKuObj.GetLuckyPoints())
	pl.SendMsg(scRingLuckyPointsChange)
	return
}
