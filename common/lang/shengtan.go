package lang

const (
	ShengTanPlayerNoInScene LangCode = ShengTanBase + iota
	ShengTanJiuNiangLimit
	ShengTanJiuNiangUseNotice
	ShengTanHpChangedNotice
	ShengTanAllianceUserNotInAlliance
	ShengTanAllianceUserNotSameAlliance
)

var (
	shengTanLangMap = map[LangCode]string{
		ShengTanPlayerNoInScene:             "玩家不在仙盟圣坛中",
		ShengTanJiuNiangLimit:               "仙盟仙酿使用次数已达上限，无法使用",
		ShengTanJiuNiangUseNotice:           "%s使用了%s，圣坛经验加成大幅上涨，赶紧前往获得经验吧！%s",
		ShengTanHpChangedNotice:             "仙盟圣坛正被入侵怪物攻击，血量仅剩%s%%，请尽快前往支援！%s",
		ShengTanAllianceUserNotInAlliance:   "您当前还没有仙盟，无法参加此活动",
		ShengTanAllianceUserNotSameAlliance: "您不是该仙盟的,无法进入仙盟圣坛",
	}
)

func init() {
	mergeLang(shengTanLangMap)
}
