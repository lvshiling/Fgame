package log

type LaBaLogReason int32

const (
	LaBaLogReasonGM LaBaLogReason = iota + 1
	LaBaLogReasonAttend
)

func (r LaBaLogReason) Reason() int32 {
	return int32(r)
}

var (
	labaLogReasonMap = map[LaBaLogReason]string{
		LaBaLogReasonGM:     "gm修改",
		LaBaLogReasonAttend: "参与元宝拉霸,当前次数:%d",
	}
)

func (r LaBaLogReason) String() string {
	return labaLogReasonMap[r]
}
