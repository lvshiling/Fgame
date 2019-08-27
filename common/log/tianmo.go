package log

type TianMoLogReason int32

const (
	TianMoLogReasonGM TianMoLogReason = iota + 1
	TianMoLogReasonAdvanced
)

func (zslr TianMoLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	tianMoLogReasonMap = map[TianMoLogReason]string{
		TianMoLogReasonGM:       "gm修改",
		TianMoLogReasonAdvanced: "天魔体进阶,进阶方式:%s",
	}
)

func (ar TianMoLogReason) String() string {
	return tianMoLogReasonMap[ar]
}
