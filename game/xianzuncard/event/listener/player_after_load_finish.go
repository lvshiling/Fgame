package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/xianzuncard/pbutil"
	playerxianzuncard "fgame/fgame/game/xianzuncard/player"
)

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}

// 玩家加载
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	//登录下发
	xianZunManager := pl.GetPlayerDataManager(playertypes.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	xianZunObjMap := xianZunManager.GetXianZunCardObjectMap()
	scXianZunCardInfo := pbutil.BuildSCXianZunCardInfo(xianZunObjMap)
	pl.SendMsg(scXianZunCardInfo)

	return
}
