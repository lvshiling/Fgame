package log

type ShieldLogReason int32

const (
	ShieldLogReasonGM ShieldLogReason = iota + 1
	ShieldLogReasonAdvanced
	ShieldLogReasonBodyAdvanced
)

func (r ShieldLogReason) Reason() int32 {
	return int32(r)
}

var (
	shieldLogReasonMap = map[ShieldLogReason]string{
		ShieldLogReasonGM:           "gm修改",
		ShieldLogReasonAdvanced:     "盾刺进阶,进阶方式:%s",
		ShieldLogReasonBodyAdvanced: "护体盾进阶,进阶方式:%s",
	}
)

func (r ShieldLogReason) String() string {
	return shieldLogReasonMap[r]
}
