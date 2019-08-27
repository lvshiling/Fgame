package log

type GoldYuanLevelLogReason int32

const (
	GoldYuanLevelLogReasonGM GoldYuanLevelLogReason = iota + 1
	GoldYuanLevelLogReasonEatEquip
)

func (zslr GoldYuanLevelLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	GoldYuanLevelLogReasonMap = map[GoldYuanLevelLogReason]string{
		GoldYuanLevelLogReasonGM:       "gm修改",
		GoldYuanLevelLogReasonEatEquip: "玩家吞噬元神金装",
	}
)

func (slr GoldYuanLevelLogReason) String() string {
	return GoldYuanLevelLogReasonMap[slr]
}
