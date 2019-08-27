package log

type GoldEquipLogReason int32

const (
	GoldEquipLogReasonGM GoldEquipLogReason = iota + 1
	GoldEquipLogReasonExtend
)

func (r GoldEquipLogReason) Reason() int32 {
	return int32(r)
}

var (
	goldEquipLogReasonMap = map[GoldEquipLogReason]string{
		GoldEquipLogReasonGM:     "gm修改",
		GoldEquipLogReasonExtend: "装备继承，继承装备id：%d,材料装备id:%d",
	}
)

func (r GoldEquipLogReason) String() string {
	return goldEquipLogReasonMap[r]
}
