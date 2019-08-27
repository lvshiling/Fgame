package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	"fgame/fgame/game/trade/pbutil"
	"fgame/fgame/game/trade/trade"
)

//上传成功
func tradeUpload(target event.EventTarget, data event.EventData) (err error) {
	refundTradeItemObj, ok := target.(*trade.TradeItemObject)
	if !ok {
		return
	}
	playerId := refundTradeItemObj.GetPlayerId()
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		return
	}
	//发送消息
	scTradeUploadItem := pbutil.BuildSCTradeUploadItem(refundTradeItemObj)
	p.SendMsg(scTradeUploadItem)
	return
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.TradeEventTypeTradeUpload, event.EventListenerFunc(tradeUpload))
}
