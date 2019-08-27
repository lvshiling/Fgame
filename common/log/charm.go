package log

type CharmLogReason int32

const (
	CharmLogReasonGM CharmLogReason = iota + 1
	CharmLogReasonGiftReward
)

func (zslr CharmLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	charmLogReasonMap = map[CharmLogReason]string{
		CharmLogReasonGM:         "gm修改",
		CharmLogReasonGiftReward: "玩家赠送奖励",
	}
)

func (slr CharmLogReason) String() string {
	return charmLogReasonMap[slr]
}
