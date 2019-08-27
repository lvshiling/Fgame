package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"

	log "github.com/Sirupsen/logrus"
)

//下单失败
func getOrderFailed(target event.EventTarget, data event.EventData) (err error) {
	playerId := target.(int64)
	chargeId := data.(int32)
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("charge:下单成功,玩家不存在")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": playerId,
			"chargeId": chargeId,
		}).Warn("charge:下单成功,下单失败")
	playerlogic.SendSystemMessage(p, lang.ChargeOrderFailed)
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeGetOrderFailed, event.EventListenerFunc(getOrderFailed))
}
