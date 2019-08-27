package lang

const (
	ReliveBaseExceedMaxTimes LangCode = ReliveBase + iota
	ReliveBaseYuanDiMaxTimes
)

var (
	reliveLangMap = map[LangCode]string{
		ReliveBaseExceedMaxTimes: "复活次数已超上限",
		ReliveBaseYuanDiMaxTimes: "原地复活已超上限",
	}
)

func init() {
	mergeLang(reliveLangMap)
}
