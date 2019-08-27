package lang

const (
	LingTongActiveSystem LangCode = LingTongBase + iota
	LingTongNoActive
	LingTongShengJiReachFull
	LingTongPeiYangReachFull
	LingTongUpstarReachFull
)

var (
	lingTongLangMap = map[LangCode]string{
		LingTongActiveSystem:     "请先激活灵童系统",
		LingTongNoActive:         "请先激活该灵童",
		LingTongShengJiReachFull: "您的该灵童升级已达满级",
		LingTongPeiYangReachFull: "您的该灵童培养已达满级",
		LingTongUpstarReachFull:  "您的该灵童升星已达满级",
	}
)

func init() {
	mergeLang(lingTongLangMap)
}
