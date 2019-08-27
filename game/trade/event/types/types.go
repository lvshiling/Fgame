package types

import (
	commonlog "fgame/fgame/common/log"
)

type TradeEventType string

const (
	TradeEventTypeTradeUploadRefund TradeEventType = "TradeUploadRefund"
	TradeEventTypeTradeUpload                      = "TradeUpload"
	TradeEventTypeTradeWithdraw                    = "TradeWithdraw"
	TradeEventTypeTradeItemRefund                  = "TradeItemRefund"
	TradeEventTypeTradeItem                        = "TradeItem"
	TradeEventTypeTradeSellItem                    = "TradeSellItem"
	TradeEventTypeTradeRecycle                     = "TradeRecycle"
)

// 上架日志
type TradeLogEventType string

const (
	EventTypeTradeLogUpload      TradeLogEventType = "TradeUploadLog"      //
	EventTypeTradeLogWithdraw                      = "TradeWithdrawLog"    //
	EventTypeTradeLogBuy                           = "TradeItemLog"        //
	EventTypeTradeLogSell                          = "SellItemLog"         //
	EventTypeTradeLogRecycleGold                   = "TradeRecycleGoldLog" //
)

type PlayerTradeLogEventData struct {
	playerId   int64
	itemId     int32
	num        int32
	gold       int32
	reason     commonlog.TradeLogReason
	reasonText string
}

func CreatePlayerTradeLogEventData(playerId int64, itemId, num, gold int32, reason commonlog.TradeLogReason, reasonText string) *PlayerTradeLogEventData {
	d := &PlayerTradeLogEventData{
		playerId:   playerId,
		itemId:     itemId,
		num:        num,
		gold:       gold,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerTradeLogEventData) GetPlayerId() int64 {
	return d.playerId
}

func (d *PlayerTradeLogEventData) GetItemId() int32 {
	return d.itemId
}

func (d *PlayerTradeLogEventData) GetNum() int32 {
	return d.num
}

func (d *PlayerTradeLogEventData) GetGold() int32 {
	return d.gold
}

func (d *PlayerTradeLogEventData) GetReason() commonlog.TradeLogReason {
	return d.reason
}

func (d *PlayerTradeLogEventData) GetReasonText() string {
	return d.reasonText
}

//
type TradeLogRecyclGoldEventData struct {
	gold       int64
	reason     commonlog.TradeLogReason
	reasonText string
}

func CreateTradeLogRecyclGoldEventData(gold int64, reason commonlog.TradeLogReason, reasonText string) *TradeLogRecyclGoldEventData {
	d := &TradeLogRecyclGoldEventData{
		gold:       gold,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *TradeLogRecyclGoldEventData) GetGold() int64 {
	return d.gold
}

func (d *TradeLogRecyclGoldEventData) GetReason() commonlog.TradeLogReason {
	return d.reason
}

func (d *TradeLogRecyclGoldEventData) GetReasonText() string {
	return d.reasonText
}
