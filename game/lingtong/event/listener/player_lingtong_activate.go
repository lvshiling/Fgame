package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerLingTongActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	lingTongId := data.(int32)
	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	manager.LingTongActivateFashionInit(lingTongId)

	//第一次激活
	lingTongObj := manager.GetLingTong()
	if lingTongObj.GetLingTongId() == 0 {
		manager.LingTongChuZhan(lingTongId)
		scLingTongChuZhan := pbutil.BuildSCLingTongChuZhan(lingTongId)
		pl.SendMsg(scLingTongChuZhan)
	}
	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeLingTongActivate, event.EventListenerFunc(playerLingTongActivate))
}
