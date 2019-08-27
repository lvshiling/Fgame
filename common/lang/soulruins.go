package lang

const (
	SoulRuinsBuyNumReachLimit LangCode = SoulRuinsBase + iota
	SoulRuinsChallengeNumNotEnough
	SoulRuinsNotReachPreLevel
	SoulRuinsRewChapterRepeatReceive
	SoulRuinsChapterStarNotEnough
	SoulRuinsNotPassedToSweep
	SoulRuinsFinishAllTitle
	SoulRuinsFinishAllContent
	SoulRuinsSoulArriveTitle
	SoulRuinsSoulArriveContent
	SoulRuinsNoExistNextLevel
)

var (
	SoulRuinsLangMap = map[LangCode]string{
		SoulRuinsBuyNumReachLimit:        "今日购买次数已达上限",
		SoulRuinsChallengeNumNotEnough:   "挑战次数不足",
		SoulRuinsNotReachPreLevel:        "请先通关前置关卡",
		SoulRuinsRewChapterRepeatReceive: "该章节星级奖励已领取过",
		SoulRuinsChapterStarNotEnough:    "您当前星数不满足领取条件",
		SoulRuinsNotPassedToSweep:        "关卡挑战成功后,才能扫荡",
		SoulRuinsFinishAllTitle:          "帝陵遗迹一键完成奖励",
		SoulRuinsFinishAllContent:        "帝陵遗迹一键完成奖励内容",
		SoulRuinsSoulArriveTitle:         "帝陵遗迹帝魂降临奖励",
		SoulRuinsSoulArriveContent:       "帝陵遗迹帝魂降临奖励内容",
		SoulRuinsNoExistNextLevel:        "当前关卡已是最高关卡,不存在下一关",
	}
)

func init() {
	mergeLang(SoulRuinsLangMap)
}
