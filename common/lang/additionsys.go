package lang

const (
	AdditionSysLevelHighest = AdditionSysBase + iota
	AdditionSysSlotNoEquip
	AdditionSysEquipmentUpgradeMax
	AdditionSysShenZhuHighest
	AdditionSysShenZhuQualityLimit
	AdditionSysShenZhuItemNoEnough
	AdditionSysAdvancedNoEnough
	AdditionSysAlreadyAwake
	AdditionSysAwakeLevelTop
)

var (
	additionSysLangMap = map[LangCode]string{
		AdditionSysLevelHighest:        "系統等级已满级",
		AdditionSysSlotNoEquip:         "装备槽没有装上装备",
		AdditionSysEquipmentUpgradeMax: "系统装备已经满阶",
		AdditionSysShenZhuHighest:      "该部位神铸等级已满级",
		AdditionSysShenZhuQualityLimit: "必须穿戴橙色品质装备才可进行神铸",
		AdditionSysShenZhuItemNoEnough: "物品不足，无法进行神铸",
		AdditionSysAdvancedNoEnough:    "系统阶数不足，无法进行觉醒",
		AdditionSysAlreadyAwake:        "系统已经觉醒",
		AdditionSysAwakeLevelTop:       "觉醒等级已满",
	}
)

func init() {
	mergeLang(additionSysLangMap)
}
