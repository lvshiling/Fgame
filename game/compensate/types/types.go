package types

// 补偿状态
type CompensateRecordSate int32

const (
	CompensateRecordSateHadGet CompensateRecordSate = iota //已获取
	CompensateRecordSateNotGet                             //不满足条件
)
