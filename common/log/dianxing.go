package log

type DianXingLogReason int32

const (
	DianXingLogReasonGM DianXingLogReason = iota + 1
	DianXingLogReasonAdvanced
	DianXingLogReasonJieFengAdvanced
)

func (dxlr DianXingLogReason) Reason() int32 {
	return int32(dxlr)
}

var (
	dianXingLogReasonMap = map[DianXingLogReason]string{
		DianXingLogReasonGM:              "gm修改",
		DianXingLogReasonAdvanced:        "点星系统升级,升级方式:%s",
		DianXingLogReasonJieFengAdvanced: "点星解封升级,升级方式:%s",
	}
)

func (dxlr DianXingLogReason) String() string {
	return dianXingLogReasonMap[dxlr]
}
