package log

type AnqiLogReason int32

const (
	AnqiLogReasonGM AnqiLogReason = iota + 1
	AnqiLogReasonAdvanced
)

func (zslr AnqiLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	anqiLogReasonMap = map[AnqiLogReason]string{
		AnqiLogReasonGM:       "gm修改",
		AnqiLogReasonAdvanced: "暗器进阶,进阶方式:%s",
	}
)

func (ar AnqiLogReason) String() string {
	return anqiLogReasonMap[ar]
}
