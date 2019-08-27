package log

type MountLogReason int32

const (
	MountLogReasonGM MountLogReason = iota + 1
	MountLogReasonAdvanced
)

func (r MountLogReason) Reason() int32 {
	return int32(r)
}

var (
	mountLogReasonMap = map[MountLogReason]string{
		MountLogReasonGM:       "gm修改",
		MountLogReasonAdvanced: "坐骑进阶,进阶方式:%s",
	}
)

func (r MountLogReason) String() string {
	return mountLogReasonMap[r]
}
