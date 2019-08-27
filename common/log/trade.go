package log

type TradeLogReason int32

const (
	TradeLogReasonUpload TradeLogReason = iota + 1
	TradeLogReasonWithdraw
	TradeLogReasonRecycleGold
	TradeLogReasonTradeItem
	TradeLogReasonSellItem
)

func (r TradeLogReason) Reason() int32 {
	return int32(r)
}

var (
	tradeLogReasonMap = map[TradeLogReason]string{
		TradeLogReasonUpload:      "上架商品",
		TradeLogReasonWithdraw:    "下架商品,是否系统操作：%v",
		TradeLogReasonRecycleGold: "交易回购池日志",
		TradeLogReasonTradeItem:   "交易所购买物品",
		TradeLogReasonSellItem:    "交易所卖出物品",
	}
)

func (ar TradeLogReason) String() string {
	return tradeLogReasonMap[ar]
}
