package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/tianmo/pbutil"
	playertianmo "fgame/fgame/game/tianmo/player"
)

//玩家充值元宝
func playerChargeGold(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	chargeGold := data.(int32)

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeTianMo) {
		return
	}

	// 记录充值额度
	tianmoManager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianmoManager.AddChargeNum(chargeGold)

	info := tianmoManager.GetTianMoInfo()
	scMsg := pbutil.BuildSCTianMoChargeGold(info.ChargeVal)
	pl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeGold))
}
