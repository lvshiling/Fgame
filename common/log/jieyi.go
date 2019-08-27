package log

type JieYiLogReason int32

const (
	JieYiLogReasonGM JieYiLogReason = iota + 1
	JieYiLogReasonDaoJuTypeChange
	JieYiLogReasonTokenTypeChange
	JieYiLogReasonTokenLevelChange
	JieYiLogReasonNameLevChange
	JieYiLogReasonShengWeiZhiChange
)

func (zslr JieYiLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	jieYiLogReasonMap = map[JieYiLogReason]string{
		JieYiLogReasonGM:                "gm修改",
		JieYiLogReasonDaoJuTypeChange:   "道具类型改变,之前类型：%s,当前类型：%s",
		JieYiLogReasonTokenTypeChange:   "信物类型变化,之前类型：%s,当前类型：%s",
		JieYiLogReasonTokenLevelChange:  "信物等级变化,之前等级：%d,当前等级：%d",
		JieYiLogReasonNameLevChange:     "威名等级变化,之前等级：%d,当前等级：%d;",
		JieYiLogReasonShengWeiZhiChange: "声威值变化,之前数量：%d,当前数量：%d",
	}
)

func (ar JieYiLogReason) String() string {
	return jieYiLogReasonMap[ar]
}
