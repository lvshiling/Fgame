package log

type QiXueLogReason int32

const (
	QiXueLogReasonAdvanced QiXueLogReason = iota + 1
	QiXueLogReasonDegrade
)

func (zslr QiXueLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	qiXueLogReasonMap = map[QiXueLogReason]string{
		QiXueLogReasonAdvanced: "泣血枪进阶,进阶方式:%s",
		QiXueLogReasonDegrade:  "泣血枪进降阶,降阶方式:被击杀",
	}
)

func (ar QiXueLogReason) String() string {
	return qiXueLogReasonMap[ar]
}
