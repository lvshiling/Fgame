package lang

const (
	SkillHasExist = SkillBase + iota
	SkillNotHas
	SkillReachLimit
	SKillNotHasUpgrade
	SkillTianFuHasedAwaken
	SkillTianFuAwakenNoAwakenParent
	SKillTianFuUpgradeNoAwaken
	SkillTianFuUpgradeFull
)

var (
	skillLangMap = map[LangCode]string{
		SkillHasExist:                   "该技能已存在,无需添加",
		SkillNotHas:                     "还没有该技能,请先获取",
		SkillReachLimit:                 "当前技能已达最高级，无法升级",
		SKillNotHasUpgrade:              "当前无可升级技能",
		SkillTianFuHasedAwaken:          "当前天赋已经觉醒过",
		SkillTianFuAwakenNoAwakenParent: "您还未觉醒【%s天赋】,无法觉醒当前天赋",
		SKillTianFuUpgradeNoAwaken:      "您的天赋还未觉醒,无法升级",
		SkillTianFuUpgradeFull:          "您的天赋等级已达满级",
	}
)

func init() {
	mergeLang(skillLangMap)
}
