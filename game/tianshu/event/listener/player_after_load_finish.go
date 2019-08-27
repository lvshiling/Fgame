package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/tianshu/pbutil"
	playertianshu "fgame/fgame/game/tianshu/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	tianshuManager := pl.GetPlayerDataManager(playertypes.PlayerTianShuDataManagerType).(*playertianshu.PlayerTianShuDataManager)

	totalChargeNum := int32(pl.GetChargeGoldNum())
	tianshuList := tianshuManager.GetTianShuAll()
	scMsg := pbutil.BuildSCTianShuInfoList(tianshuList, totalChargeNum)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
