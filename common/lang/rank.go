package lang

const (
	RankLoginFirst LangCode = RankBase + iota
	RankLoginSecond
	RankLoginThird
)

var (
	rankLangMap = map[LangCode]string{
		RankLoginFirst:  "凌驾于众人之上的战力榜第一:%s降临修仙世界。霎时间日月无光,星辰失色，强大的威压令所有的修仙者都无法呼吸！",
		RankLoginSecond: "战力榜第二:%s降临修仙世界，众人为他送上喝彩与掌声！",
		RankLoginThird:  "战力榜第三:%s降临修仙世界，众人为他送上喝彩与掌声！",
	}
)

func init() {
	mergeLang(rankLangMap)
}
