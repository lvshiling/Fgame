package types

type GuaJiGlobalType int32

const (
	GuaJiGlobalTypeAdvanceAutoBuy GuaJiGlobalType = iota //挂机进阶自动购买
	GuaJiGlobalTypeBagRemainSlots                        //需要清空的包剩余位置
	GuaJiAutoBuyBagLevel                                 //自动购买槽位等级
)
