package log

type MarryLogReason int32

const (
	MarryLogReasonGM MarryLogReason = iota + 1
	MarryLogReasonDevelop
	MarryLogReasonDevelopExpToLevel
	MarryLogReasonDevelopExpByItem
)

func (zslr MarryLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	marryLogReasonMap = map[MarryLogReason]string{
		MarryLogReasonGM:                "gm修改",
		MarryLogReasonDevelop:           "表白等级升级",
		MarryLogReasonDevelopExpToLevel: "表白经验升级等级,当前等级[%d],之前等级[%d]",
		MarryLogReasonDevelopExpByItem:  "表白经验使用物品[%d],数量[%d]",
	}
)

func (ar MarryLogReason) String() string {
	return marryLogReasonMap[ar]
}
