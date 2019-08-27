package types

type MajorPairCodeType int32

const (
	//成功
	MajorPairCodeTypeSucess MajorPairCodeType = 1
	//配偶已决策
	MajorPairCodeTypeDeal MajorPairCodeType = 2
	//邀请已取消
	MajorPairCodeTypeCancle MajorPairCodeType = 2
)
