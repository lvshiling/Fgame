package lang

const (
	CrossFailed LangCode = CrossBase + iota
)

var (
	crossLangMap = map[LangCode]string{

		CrossFailed: "跨服失败",
	}
)

func init() {
	mergeLang(crossLangMap)
}
