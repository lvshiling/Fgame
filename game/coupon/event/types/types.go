package types

type CouponEventType string

const (
	CouponEventTypeExchangeFinish CouponEventType = "ExchangeFinish" //充值元宝
	CouponEventTypeExchangeFailed                 = "ExchangeFailed"
)
