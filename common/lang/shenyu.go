package lang

const (
	ShenYuFirstRoundEnd LangCode = ShenYuBase + iota
	ShenYuNotShenYu
	ShenYuHadAttend
	ShenYuLuckyRewMailTitle
	ShenYuLuckyRewMailContent
	ShenYuRankMailTitle
	ShenYuRankMailContent
)

var (
	shenYuLangMap = map[LangCode]string{
		ShenYuFirstRoundEnd:       "神域之战第一轮已结束",
		ShenYuNotShenYu:           "不是神域之战",
		ShenYuHadAttend:           "已参与过神域之战",
		ShenYuLuckyRewMailTitle:   "神域幸运奖",
		ShenYuLuckyRewMailContent: "恭喜您成功在神域之战中获得【幸运奖】，奖励如下，请查收！",
		ShenYuRankMailTitle:       "神域排行版",
		ShenYuRankMailContent:     "恭喜您成功在神域之战中获得【%d】名，奖励如下，请查收",
	}
)

func init() {
	mergeLang(shenYuLangMap)
}
