package lang

const (
	QuizAnswerTimeLimit = QuizBase + iota
	QuizAnswerAlreadyGet
	QuizAnswerRewTitle
	QuizAnswerRewContent
)

var (
	quizLangMap = map[LangCode]string{
		QuizAnswerTimeLimit:  "超过答题时间",
		QuizAnswerAlreadyGet: "已经答过题目了",
		QuizAnswerRewTitle:   "仙尊问答奖励",
		QuizAnswerRewContent: "答题奖励，背包不足",
	}
)

func init() {
	mergeLang(quizLangMap)
}
