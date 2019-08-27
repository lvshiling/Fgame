package log

type EquipBaoKuLogReason int32

const (
	EquipBaoKuLogReasonGM EquipBaoKuLogReason = iota + 1
	EquipBaoKuLogReasonAttendPointsChange
	EquipBaoKuLogReasonLuckyPointsChange
	EquipBaoKuLogLuckyPointRewContent
)

func (zslr EquipBaoKuLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	equipBaoKuLogReasonMap = map[EquipBaoKuLogReason]string{
		EquipBaoKuLogReasonGM:                 "gm修改",
		EquipBaoKuLogReasonAttendPointsChange: "%s积分变化,变化方式:%s",
		EquipBaoKuLogReasonLuckyPointsChange:  "%s幸运值变化,变化方式:%s",
		EquipBaoKuLogLuckyPointRewContent:     "物品id %d,数量 %d;",
	}
)

func (ar EquipBaoKuLogReason) String() string {
	return equipBaoKuLogReasonMap[ar]
}
