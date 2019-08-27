package lang

const (
	Unknown LangCode = UnknownBase + iota + 1
)

var (
	unknownLangMap = map[LangCode]string{
		Unknown: "不知道的错误",
	}
)

func init() {
	mergeLang(unknownLangMap)
}
