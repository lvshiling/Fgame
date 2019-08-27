package log

type ZhuanShengLogReason int32

const (
	ZhuanShengLogReasonGM ZhuanShengLogReason = iota + 1
	ZhuanShengLogReasonQuestReward
	ZhuanShengLogReasonGoldEquip
	ZhuanShengLogReasonGuaJi
)

func (zslr ZhuanShengLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	zhuanShengLogReasonMap = map[ZhuanShengLogReason]string{
		ZhuanShengLogReasonGM:          "gm修改",
		ZhuanShengLogReasonQuestReward: "任务完成奖励,任务id:%d",
		ZhuanShengLogReasonGoldEquip:   "元神金装转生请求",
		ZhuanShengLogReasonGuaJi:       "内挂转生",
	}
)

func (slr ZhuanShengLogReason) String() string {
	return zhuanShengLogReasonMap[slr]
}
