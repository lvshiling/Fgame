package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家交易下架日志
func playerTradeWithdrawLog(target event.EventTarget, data event.EventData) (err error) {

	eventData, ok := data.(*tradeeventtypes.PlayerTradeLogEventData)
	if !ok {
		return
	}

	tradeDrawLog := &logmodel.PlayerTradeWithdraw{}
	tradeDrawLog.PlayerTradeLogMsg = gamelog.PlayerTradeLogMsgFromTradeObject(eventData.GetPlayerId())
	tradeDrawLog.ItemId = eventData.GetItemId()
	tradeDrawLog.ItemNum = eventData.GetNum()
	tradeDrawLog.Gold = eventData.GetGold()
	tradeDrawLog.Reason = int32(eventData.GetReason())
	tradeDrawLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(tradeDrawLog)
	return
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.EventTypeTradeLogWithdraw, event.EventListenerFunc(playerTradeWithdrawLog))
}
