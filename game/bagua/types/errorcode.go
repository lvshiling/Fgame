package types

type BaGuaPairCodeType int32

const (
	//成功
	BaGuaPairCodeTypeSucess BaGuaPairCodeType = 1
	//配偶已决策
	BaGuaPairCodeTypeDeal BaGuaPairCodeType = 2
	//邀请已取消
	BaGuaPairCodeTypeCancle BaGuaPairCodeType = 3
)
