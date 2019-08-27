package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家交易卖出日志
func playerTradeSellLog(target event.EventTarget, data event.EventData) (err error) {

	eventData, ok := data.(*tradeeventtypes.PlayerTradeLogEventData)
	if !ok {
		return
	}

	tradeSellLog := &logmodel.PlayerTradeSell{}
	tradeSellLog.PlayerTradeLogMsg = gamelog.PlayerTradeLogMsgFromTradeObject(eventData.GetPlayerId())
	tradeSellLog.ItemId = eventData.GetItemId()
	tradeSellLog.ItemNum = eventData.GetNum()
	tradeSellLog.Gold = eventData.GetGold()
	tradeSellLog.Reason = int32(eventData.GetReason())
	tradeSellLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(tradeSellLog)
	return
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.EventTypeTradeLogSell, event.EventListenerFunc(playerTradeSellLog))
}
