package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shop/pbutil"
	playershop "fgame/fgame/game/shop/player"
)

//玩家登录成功后下发,商铺每日限购道具,已购买次数
func playerShopAfterLogin(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	shopManager := pl.GetPlayerDataManager(playertypes.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	shops := shopManager.GetShopBuyAll()
	scShopLimit := pbutil.BuildSCShopLimit(shops)
	err = pl.SendMsg(scShopLimit)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerShopAfterLogin))
}
