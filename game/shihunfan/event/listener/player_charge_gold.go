package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shihunfan/pbutil"
	playershihunfan "fgame/fgame/game/shihunfan/player"
)

//玩家充值元宝
func playerChargeGold(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	//功能要开启
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShiHunFan) {
		return
	}
	chargeGold := data.(int32)

	manager := pl.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	manager.ShiHunFanCharge(chargeGold)
	shiHunFanInfo := manager.GetShiHunFanInfo()

	scMsg := pbutil.BuildSCShiHunFanChargeVary(shiHunFanInfo.ChargeVal)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeGold))
}
