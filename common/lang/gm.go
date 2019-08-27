package lang

const (
	GMFormatWrong LangCode = GMBase + iota
	GMCanNotUse
)

var (
	gmLangMap = map[LangCode]string{
		GMFormatWrong: "gm命令格式不对",
		GMCanNotUse:   "不能使用该gm命令",
	}
)

func init() {
	mergeLang(gmLangMap)
}
