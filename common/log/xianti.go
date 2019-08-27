package log

type XianTiLogReason int32

const (
	XianTiLogReasonGM XianTiLogReason = iota + 1
	XianTiLogReasonAdvanced
)

func (r XianTiLogReason) Reason() int32 {
	return int32(r)
}

var (
	xianTiLogReasonMap = map[XianTiLogReason]string{
		XianTiLogReasonGM:       "gm修改",
		XianTiLogReasonAdvanced: "仙体进阶,进阶方式:%s",
	}
)

func (r XianTiLogReason) String() string {
	return xianTiLogReasonMap[r]
}
