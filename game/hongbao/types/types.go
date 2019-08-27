package types

type HongBaoResultType int32

const (
	//0抢完
	HongBaoResultTypeFinish HongBaoResultType = iota
	//1抢过
	HongBaoResultTypeSnatched
	//2过期
	HongBaoResultTypeEndTime
	//3成功
	HongBaoResultTypeSucceed
)
