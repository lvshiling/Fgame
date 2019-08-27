package types

// 连续充值奖励类型
type FeedbackCycleRewType int32

const (
	FeedbackCycleRewTypeSingleDay FeedbackCycleRewType = iota //每日奖励
	FeedbackCycleRewTypeCountDay                              //累计奖励
)

func (t FeedbackCycleRewType) Valid() bool {
	switch t {
	case FeedbackCycleRewTypeSingleDay,
		FeedbackCycleRewTypeCountDay:
		return true
	default:
		return false
	}
}
