package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	shopeventtypes "fgame/fgame/game/shop/event/types"
	"fgame/fgame/game/shop/shop"
	logmodel "fgame/fgame/logserver/model"
)

//玩家商店改物品日志
func playerShopBuyItemLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*shopeventtypes.PlayerShopBuyItemLogEventData)
	if !ok {
		return
	}

	logPlayerShopBuy := &logmodel.PlayerShopBuy{}
	logPlayerShopBuy.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logPlayerShopBuy.ShopId = eventData.GetShopId()
	shopTemplate := shop.GetShopService().GetShopTemplate(int(eventData.GetShopId()))
	shopName := ""
	if shopTemplate != nil {
		shopName = shopTemplate.Name
	}
	logPlayerShopBuy.ShopName = shopName
	logPlayerShopBuy.BuyNum = eventData.GetBuyNum()
	logPlayerShopBuy.CostMoney = eventData.GetCostMoney()
	logPlayerShopBuy.Reason = int32(eventData.GetReason())
	logPlayerShopBuy.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logPlayerShopBuy)
	return
}

func init() {
	gameevent.AddEventListener(shopeventtypes.EventTypeShopBuyItemLog, event.EventListenerFunc(playerShopBuyItemLog))
}
