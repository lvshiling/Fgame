package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/charge/charge"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	"fgame/fgame/game/charge/pbutil"
	chargetemplate "fgame/fgame/game/charge/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//下单成功
func getOrderFinish(target event.EventTarget, data event.EventData) (err error) {
	playerId := target.(int64)
	eventData := data.(*charge.GetOrderFinishEventData)
	orderId := eventData.GetOrderId()
	chargeId := eventData.GetChargeId()
	notifyUrl := eventData.GetNotifyUrl()
	serverId := eventData.GetServerId()
	platformUserId := eventData.GetPlatformUserId()

	money := eventData.GetMoney()
	playerName := eventData.GetName()

	serverName := fmt.Sprintf("%d", serverId)
	sdkOrderId := eventData.GetSdkOrderId()
	extension := eventData.GetExtension()

	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"orderId":  orderId,
				"chargeId": chargeId,
			}).Warn("charge:下单成功,玩家不存在")
		return
	}
	chargeTemplate := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	if chargeTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"orderId":  orderId,
				"chargeId": chargeId,
			}).Warn("charge:下单成功,模板不存在")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":       playerId,
			"orderId":        orderId,
			"chargeId":       chargeId,
			"platformUserId": platformUserId,
			"money":          money,
		}).Info("charge:下单成功")
	scChargeOrder := pbutil.BuildSCChargeOrder(
		chargeId,
		chargeTemplate.Type,
		orderId,
		notifyUrl,
		sdkOrderId,
		platformUserId,
		money,
		playerId,
		playerName,
		serverId,
		serverName,
		extension)
	p.SendMsg(scChargeOrder)

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeGetOrderFinish, event.EventListenerFunc(getOrderFinish))
}
