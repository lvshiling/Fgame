package log

type AdditionSysLogReason int32

const (
	AdditionSysLogReasonGM AdditionSysLogReason = iota + 1
	AdditionSysLogReasonStrengthenUpgrade
	AdditionSysLogReasonStrengthenBackLev
	AdditionSysLogReasonShengJi
	AdditionSysLogReasonShenZhu
	AdditionSysLogReasonTongLing
	AdditionSysLogReasonAwake
)

func (zslr AdditionSysLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	additionSysLogReasonMap = map[AdditionSysLogReason]string{
		AdditionSysLogReasonGM:                "gm修改",
		AdditionSysLogReasonStrengthenUpgrade: "系统%s部位%s升级,升级方式强化",
		AdditionSysLogReasonStrengthenBackLev: "系统%s部位%s降级,降级方式强化",
		AdditionSysLogReasonShengJi:           "系统%s升级,升级方式:%s",
		AdditionSysLogReasonShenZhu:           "系统%s部位%s神铸,升级方式:%s",
		AdditionSysLogReasonTongLing:          "系统%s通灵升级,升级方式:%s",
		AdditionSysLogReasonAwake:             "系统%s觉醒,觉醒方式:%s",
	}
)

func (ar AdditionSysLogReason) String() string {
	return additionSysLogReasonMap[ar]
}
