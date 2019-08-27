package types

type OrderStatus int32

const (
	OrderStatusInit OrderStatus = iota
	OrderStatusPay
	OrderStatusFinish
	OrderStatusCancel
	OrderStatusFail
)

var (
	orderStatusMap = map[OrderStatus]string{
		OrderStatusInit:   "初始化",
		OrderStatusPay:    "充值",
		OrderStatusFinish: "发货成功",
		OrderStatusCancel: "取消",
		OrderStatusFail:   "充值失败",
	}
)

func (s OrderStatus) String() string {
	return orderStatusMap[s]
}
