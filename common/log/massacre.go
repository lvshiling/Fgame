package log

type MassacreLogReason int32

const (
	MassacreLogReasonGM MassacreLogReason = iota + 1
	MassacreLogReasonAdvanced
	MassacreLogReasonDegrade
)

func (zslr MassacreLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	massacreLogReasonMap = map[MassacreLogReason]string{
		MassacreLogReasonGM:       "gm修改",
		MassacreLogReasonAdvanced: "戮仙刃进阶,进阶方式:%s",
		MassacreLogReasonDegrade:  "戮仙刃进降阶,降阶方式:被击杀",
	}
)

func (ar MassacreLogReason) String() string {
	return massacreLogReasonMap[ar]
}
