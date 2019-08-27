package log

type ShenfaLogReason int32

const (
	ShenfaLogReasonGM ShenfaLogReason = iota + 1
	ShenfaLogReasonAdvanced
)

func (r ShenfaLogReason) Reason() int32 {
	return int32(r)
}

var (
	shenfaLogReasonMap = map[ShenfaLogReason]string{
		ShenfaLogReasonGM:       "gm修改",
		ShenfaLogReasonAdvanced: "身法进阶,进阶方式:%s",
	}
)

func (r ShenfaLogReason) String() string {
	return shenfaLogReasonMap[r]
}
