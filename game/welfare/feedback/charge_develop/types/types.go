package types

//养鸡培养奖励类型
type FeedbackDevelopRewType int32

const (
	FeedbackDevelopRewTypeSingleDay FeedbackDevelopRewType = iota //每日奖励
	FeedbackDevelopRewTypeCountDay                                //累计奖励
	FeedbackDevelopRewTypeCondition                               //开启条件
)

func (t FeedbackDevelopRewType) Valid() bool {
	switch t {
	case FeedbackDevelopRewTypeSingleDay,
		FeedbackDevelopRewTypeCountDay,
		FeedbackDevelopRewTypeCondition:
		return true
	default:
		return false
	}
}
