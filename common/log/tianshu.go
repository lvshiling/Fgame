package log

type TianShuLogReason int32

const (
	TianShuLogReasonGM TianShuLogReason = iota + 1
	TianShuLogReasonActivate
	TianShuLogReasonUplevel
)

func (zslr TianShuLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	tianshuLogReasonMap = map[TianShuLogReason]string{
		TianShuLogReasonGM:       "gm修改",
		TianShuLogReasonActivate: "天书激活,天书类型:%d",
		TianShuLogReasonUplevel:  "天书升级,天书类型:%d",
	}
)

func (ar TianShuLogReason) String() string {
	return tianshuLogReasonMap[ar]
}
