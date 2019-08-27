package log

type XianFuLogReason int32

const (
	XianFuLogReasonGM XianFuLogReason = iota + 1
	XianFuLogReasonUpgrade
)

func (zslr XianFuLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	xianfuLogReasonMap = map[XianFuLogReason]string{
		XianFuLogReasonGM:      "gm修改",
		XianFuLogReasonUpgrade: "仙府升级,仙府类型:%d",
	}
)

func (ar XianFuLogReason) String() string {
	return xianfuLogReasonMap[ar]
}
