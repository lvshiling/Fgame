package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/zhenfa/pbutil"
	playerzhenfa "fgame/fgame/game/zhenfa/player"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerZhenFaDataManagerType).(*playerzhenfa.PlayerZhenFaDataManager)
	//阵法
	zhenFaMap := manager.GetZhenFaMap()
	scZhenFaGet := pbutil.BuildSCZhenFaGet(zhenFaMap)
	p.SendMsg(scZhenFaGet)
	//阵旗
	zhenQiMap := manager.GetZhenQiMap()
	scZhenQiGet := pbutil.BuildSCZhenQiGet(zhenQiMap)
	p.SendMsg(scZhenQiGet)
	//阵旗仙火
	zhenQiXianHuoMap := manager.GetZhenQiXianHuoMap()
	scZhenQiXianHuoGet := pbutil.BuildSCZhenQiXianHuoGet(zhenQiXianHuoMap)
	p.SendMsg(scZhenQiXianHuoGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
