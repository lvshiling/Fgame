package log

type ShenQiLogReason int32

const (
	ShenQiLogReasonGM ShenQiLogReason = iota + 1
	ShenQiLogReasonRelatedUpLevel
)

func (zslr ShenQiLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	shenQiLogReasonMap = map[ShenQiLogReason]string{
		ShenQiLogReasonGM:             "gm修改",
		ShenQiLogReasonRelatedUpLevel: "神器相关升级,神器类型:%s 部位:%s, 进阶方式:%s",
	}
)

func (ar ShenQiLogReason) String() string {
	return shenQiLogReasonMap[ar]
}
