package log

type LingyuLogReason int32

const (
	LingyuLogReasonGM LingyuLogReason = iota + 1
	LingyuLogReasonAdvanced
)

func (r LingyuLogReason) Reason() int32 {
	return int32(r)
}

var (
	lingyuLogReasonMap = map[LingyuLogReason]string{
		LingyuLogReasonGM:       "gm修改",
		LingyuLogReasonAdvanced: "领域进阶,进阶方式:%s",
	}
)

func (r LingyuLogReason) String() string {
	return lingyuLogReasonMap[r]
}
