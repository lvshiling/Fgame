package types

type TradeStatus int32

const (
	TradeStatusInit       TradeStatus = iota //上架
	TradeStatusSell                          //卖出
	TradeStatusWithdraw                      //下架
	TradeStatusSellNotice                    //卖出通知

)
