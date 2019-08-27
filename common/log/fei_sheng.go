package log

type FeiShengLogReason int32

const (
	FeiShengLogReasonGM FeiShengLogReason = iota + 1
	FeiShengLogReasonDuJi
)

func (r FeiShengLogReason) Reason() int32 {
	return int32(r)
}

var (
	feiShengLogReasonMap = map[FeiShengLogReason]string{
		FeiShengLogReasonGM:   "gm修改",
		FeiShengLogReasonDuJi: "飞升渡劫，结果:%v",
	}
)

func (r FeiShengLogReason) String() string {
	return feiShengLogReasonMap[r]
}
