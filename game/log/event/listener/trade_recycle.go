package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//交易自定义回购池日志
func tradeCustomRecycleGoldLog(target event.EventTarget, data event.EventData) (err error) {

	eventData, ok := data.(*tradeeventtypes.TradeLogRecyclGoldEventData)
	if !ok {
		return
	}

	systemTradeLog := &logmodel.SystemTradeRecycle{}
	systemTradeLog.SystemLogMsg = gamelog.SystemLogMsg()
	systemTradeLog.RecycleGold = eventData.GetGold()
	systemTradeLog.Reason = int32(eventData.GetReason())
	systemTradeLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(systemTradeLog)
	return
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.EventTypeTradeLogRecycleGold, event.EventListenerFunc(tradeCustomRecycleGoldLog))
}
