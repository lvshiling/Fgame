package lang

const (
	QiXueAdanvacedReachedLimit = QiXueBase + iota
	QiXueAdanvacedShaLuNotEnough
)

var (
	qiXueLangMap = map[LangCode]string{
		QiXueAdanvacedReachedLimit:   "泣血枪已达最高阶",
		QiXueAdanvacedShaLuNotEnough: "杀戮之心不足",
	}
)

func init() {
	mergeLang(qiXueLangMap)
}
