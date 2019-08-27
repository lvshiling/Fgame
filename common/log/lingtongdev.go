package log

type LingTongDevLogReason int32

const (
	LingTongDevLogReasonGM LingTongDevLogReason = iota + 1
	LingTongDevLogReasonAdvanced
)

func (r LingTongDevLogReason) Reason() int32 {
	return int32(r)
}

var (
	lingTongDevLogReasonMap = map[LingTongDevLogReason]string{
		LingTongDevLogReasonGM:       "gm修改",
		LingTongDevLogReasonAdvanced: "灵童%s进阶,进阶方式:%s",
	}
)

func (r LingTongDevLogReason) String() string {
	return lingTongDevLogReasonMap[r]
}
