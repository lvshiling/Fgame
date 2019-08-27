package log

type FaBaoLogReason int32

const (
	FaBaoLogReasonGM FaBaoLogReason = iota + 1
	FaBaoLogReasonAdvanced
)

func (r FaBaoLogReason) Reason() int32 {
	return int32(r)
}

var (
	faBaoLogReasonMap = map[FaBaoLogReason]string{
		FaBaoLogReasonGM:       "gm修改",
		FaBaoLogReasonAdvanced: "法宝进阶,进阶方式:%s",
	}
)

func (r FaBaoLogReason) String() string {
	return faBaoLogReasonMap[r]
}
