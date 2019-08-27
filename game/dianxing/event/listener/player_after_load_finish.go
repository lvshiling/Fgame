package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/dianxing/pbutil"
	playerdianxing "fgame/fgame/game/dianxing/player"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}

	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeDianXing) {
		return
	}

	//点星系统信息
	dianxingManager := p.GetPlayerDataManager(playertypes.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	dianXingObject := dianxingManager.GetDianXingObject()
	scDianXingGet := pbutil.BuildSCDianXingGet(dianXingObject)
	p.SendMsg(scDianXingGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
