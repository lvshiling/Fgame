package log

type WingLogReason int32

const (
	WingLogReasonGM WingLogReason = iota + 1
	WingLogReasonAdvanced
	WingLogReasonFeatherAdvanced
)

func (r WingLogReason) Reason() int32 {
	return int32(r)
}

var (
	wingLogReasonMap = map[WingLogReason]string{
		WingLogReasonGM:              "gm修改",
		WingLogReasonAdvanced:        "战翼进阶,进阶方式:%s",
		WingLogReasonFeatherAdvanced: "仙羽进阶,进阶方式:%s",
	}
)

func (r WingLogReason) String() string {
	return wingLogReasonMap[r]
}
