package log

type XueDunLogReason int32

const (
	XueDunLogReasonGM XueDunLogReason = iota + 1
	XueDunLogReasonUpgrade
)

func (r XueDunLogReason) Reason() int32 {
	return int32(r)
}

var (
	xueDunLogReasonMap = map[XueDunLogReason]string{
		XueDunLogReasonGM:      "gm修改",
		XueDunLogReasonUpgrade: "血盾升阶",
	}
)

func (r XueDunLogReason) String() string {
	return xueDunLogReasonMap[r]
}
