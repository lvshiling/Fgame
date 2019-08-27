package log

type YaoPaiLogReason int32

const (
	YaoPaiLogReasonGM YaoPaiLogReason = iota + 1
	YaoPaiLogReasonPlayerKilled
	YaoPaiLogReasonExchange
	YaoPaiLogReasonConvert
)

func (glr YaoPaiLogReason) Reason() int32 {
	return int32(glr)
}

var (
	yaoPaiLogReasonMap = map[YaoPaiLogReason]string{
		YaoPaiLogReasonGM:           "gm修改",
		YaoPaiLogReasonPlayerKilled: "玩家击杀",
		YaoPaiLogReasonExchange:     "兑换:%d",
	}
)

func (glr YaoPaiLogReason) String() string {
	return yaoPaiLogReasonMap[glr]
}
