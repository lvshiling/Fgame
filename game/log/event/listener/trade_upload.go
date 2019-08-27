package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家交易上架日志
func playerTradeUploadLog(target event.EventTarget, data event.EventData) (err error) {

	eventData, ok := data.(*tradeeventtypes.PlayerTradeLogEventData)
	if !ok {
		return
	}

	tradeUploadLog := &logmodel.PlayerTradeUpload{}
	tradeUploadLog.PlayerTradeLogMsg = gamelog.PlayerTradeLogMsgFromTradeObject(eventData.GetPlayerId())
	tradeUploadLog.ItemId = eventData.GetItemId()
	tradeUploadLog.ItemNum = eventData.GetNum()
	tradeUploadLog.Gold = eventData.GetGold()
	tradeUploadLog.Reason = int32(eventData.GetReason())
	tradeUploadLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(tradeUploadLog)
	return
}

func init() {
	gameevent.AddEventListener(tradeeventtypes.EventTypeTradeLogUpload, event.EventListenerFunc(playerTradeUploadLog))
}
