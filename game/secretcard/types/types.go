package types

const (
	//窥探任务个数
	SpyTotalNum = int32(3)
	//低级池任务个数
	SpyLowPoolNum = int32(2)
)

type SecretCardPoolType int32

const (
	//低级池
	SecretCardPoolTypeNormal SecretCardPoolType = 1 + iota
	//高级池
	SecretCardPoolTypeSenior
	//轮询池
	SecretCardPoolTypePoll
)

func (scpt SecretCardPoolType) Valid() bool {
	switch scpt {
	case SecretCardPoolTypeNormal,
		SecretCardPoolTypeSenior,
		SecretCardPoolTypePoll:
		return true
	}
	return false
}
