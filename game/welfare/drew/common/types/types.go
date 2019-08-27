package types

//抽奖活动类型
type LuckyDrewAttendType int32

const (
	defaultNum = int32(1)
	batchNum   = int32(10)
)

const (
	LuckyDrewTypeOnce  LuckyDrewAttendType = iota //一次
	LuckyDrewTypeBatch                            //十连
)

func (t LuckyDrewAttendType) Valid() bool {
	switch t {
	case LuckyDrewTypeOnce,
		LuckyDrewTypeBatch:
		return true
	default:
		return false
	}
}

func (t LuckyDrewAttendType) GetAttendNum() int32 {
	switch t {
	case LuckyDrewTypeOnce:
		return defaultNum
	case LuckyDrewTypeBatch:
		return batchNum
	default:
		return 0
	}
}
