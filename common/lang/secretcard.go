package lang

const (
	SecretCardNoEnough LangCode = SecretCardBase + iota
	SecretCardStarNoEnough
	SecretCardNoQuest
	SecretCardFinishAll
	SecretCardHasExecute
	SecretCardFinishAllTitle
	SecretCardFinishAllContent
	ScecretCardFinishTitle
	ScecretCardFinishContent
	SecretImmediateFinish
	SecretCardImmediateFinishTitle
	SecretCardImmediateFinishContent
	SecretCardFinishAllVipNoEnough
)

var (
	secretCardLangMap = map[LangCode]string{
		SecretCardNoEnough:               "天机牌次数已用完",
		SecretCardStarNoEnough:           "条件不足或已领取完",
		SecretCardNoQuest:                "任务不存在",
		SecretCardFinishAll:              "今日任务已全部完成",
		SecretCardHasExecute:             "有天机任务在执行",
		SecretCardFinishAllTitle:         "天机牌一键完成奖励",
		SecretCardFinishAllContent:       "天机牌一键完成奖励内容",
		ScecretCardFinishTitle:           "天机牌完成%s奖励",
		ScecretCardFinishContent:         "天机牌完成任务奖励内容",
		SecretImmediateFinish:            "请先接取任务",
		SecretCardImmediateFinishTitle:   "天机牌直接完成奖励",
		SecretCardImmediateFinishContent: "天机牌直接完成奖励内容",
		SecretCardFinishAllVipNoEnough:   "玩家达到VIP%s以后,天机牌任务才能一键完成",
	}
)

func init() {
	mergeLang(secretCardLangMap)
}
