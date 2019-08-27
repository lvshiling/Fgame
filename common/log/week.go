package log

type WeekLogReason int32

const (
	WeekLogReasonBuy WeekLogReason = iota + 1
)

func (zslr WeekLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	weekLogReasonMap = map[WeekLogReason]string{
		WeekLogReasonBuy: "周卡购买,周卡类型:%s",
	}
)

func (ar WeekLogReason) String() string {
	return weekLogReasonMap[ar]
}
