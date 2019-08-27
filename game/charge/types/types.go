package types

type OrderStatus int32

const (
	OrderStatusInit   OrderStatus = iota //初始化
	OrderStatusFinish                    //完成
)
