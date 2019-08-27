package log

type VipLogReason int32

const (
	VipLogReasonGM VipLogReason = iota + 1
	VipLogReasonUplevel
)

func (zslr VipLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	vipLogReasonMap = map[VipLogReason]string{
		VipLogReasonGM:      "gm修改",
		VipLogReasonUplevel: "vip升级",
	}
)

func (ar VipLogReason) String() string {
	return vipLogReasonMap[ar]
}
