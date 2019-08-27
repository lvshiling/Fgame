package log

type ShiHunFanLogReason int32

const (
	ShiHunFanLogReasonGM ShiHunFanLogReason = iota + 1
	ShiHunFanLogReasonAdvanced
)

func (zslr ShiHunFanLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	shiHunFanLogReasonMap = map[ShiHunFanLogReason]string{
		ShiHunFanLogReasonGM:       "gm修改",
		ShiHunFanLogReasonAdvanced: "噬魂幡进阶,进阶方式:%s",
	}
)

func (ar ShiHunFanLogReason) String() string {
	return shiHunFanLogReasonMap[ar]
}
