package types

type TradeStatus int32

const (
	TradeStatusInit        TradeStatus = iota //初始化
	TradeStatusUpload                         //上架
	TradeStatusWithDrawing                    //下架中
	TradeStatusWithDraw                       //下架
	TradeStatusSold                           //卖出
	TradeStatusRefund                         //退还
	TradeStatusSoldNotice                     //卖出通知
)

type TradeOrderStatus int32

const (
	TradeOrderStatusInit   TradeOrderStatus = iota //支付
	TradeOrderStatusFinish                         //完成
	TradeOrderStatusRefund                         //退还
	TradeOrderStatusEnd                            //发货成功
)

type TradeLogType int32

const (
	TradeLogTypeSell TradeLogType = iota
	TradeLogTypeBuy
)
