package lang

const (
	SystemSkillRepeatActive LangCode = SystemSkillBase + iota
	SystemSkillActiveNoNumber
	SystemSkillNotActiveNotUpgrade
	SystemSkillReacheFullUpgrade
	SystemSkillActiveNotEnoughEquipNum
)

var (
	systemSkillLangMap = map[LangCode]string{
		SystemSkillRepeatActive:            "该技能已激活,无需激活",
		SystemSkillActiveNoNumber:          "系统阶数不足,无法激活",
		SystemSkillNotActiveNotUpgrade:     "未激活的技能,无法升级",
		SystemSkillReacheFullUpgrade:       "技能已达最高级",
		SystemSkillActiveNotEnoughEquipNum: "装备数量不足，无法激活！",
	}
)

func init() {
	mergeLang(systemSkillLangMap)
}
