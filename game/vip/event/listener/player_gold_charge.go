package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	viplogic "fgame/fgame/game/vip/logic"
	playervip "fgame/fgame/game/vip/player"
)

//玩家充值元宝
func playerChargeFirst(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}

	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	vipManager.UpdateChargeNum(int64(goldNum))

	viplogic.VipInfoNotice(pl)
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeFirst))
}
