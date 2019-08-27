package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/coupon/coupon"
	couponeventtypes "fgame/fgame/game/coupon/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"

	log "github.com/Sirupsen/logrus"
)

//TODO 补充错误信息
//下单失败
func exchangeCouponFailed(target event.EventTarget, data event.EventData) (err error) {
	playerId := target.(int64)
	// eventData := data.(*coupon.ExchangeFailedEventData)
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("coupon:兑换兑换码,兑换失败")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": playerId,
		}).Warn("coupon:兑换兑换码,兑换失败")
	eventData, ok := data.(*coupon.ExchangeFailedEventData)
	if !ok {
		return
	}

	playerlogic.SendSystemContentMessage(p, eventData.GetMsg())
	return
}

func init() {
	gameevent.AddEventListener(couponeventtypes.CouponEventTypeExchangeFailed, event.EventListenerFunc(exchangeCouponFailed))
}
