package types

type ExchangeStatus int32

const (
	ExchangeStatusInit    ExchangeStatus = iota //初始化
	ExchangeStatusExpired                       //过期
	ExchangeStatusFinish                        //完成
	ExchangeStatusNotify                        //通知
)
