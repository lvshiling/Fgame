package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家交易购买日志
func playerTradeBuyLog(target event.EventTarget, data event.EventData) (err error) {

	eventData, ok := data.(*tradeeventtypes.PlayerTradeLogEventData)
	if !ok {
		return
	}

	tradeBuyLog := &logmodel.PlayerTradeBuy{}
	tradeBuyLog.PlayerTradeLogMsg = gamelog.PlayerTradeLogMsgFromTradeObject(eventData.GetPlayerId())
	tradeBuyLog.ItemId = eventData.GetItemId()
	tradeBuyLog.ItemNum = eventData.GetNum()
	tradeBuyLog.Gold = eventData.GetGold()
	tradeBuyLog.Reason = int32(eventData.GetReason())
	tradeBuyLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(tradeBuyLog)
	return
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.EventTypeTradeLogBuy, event.EventListenerFunc(playerTradeBuyLog))
}
