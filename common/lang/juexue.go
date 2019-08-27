package lang

const (
	JueXueRepeatActive LangCode = JueXueBase + iota
	JueXueNotActiveNotUpgrade
	JueXueReacheFullUpgrade
	JueXueNotActiveNotInsight
	JueXueInsightNotReach
	JueXueUseNotActive
	JueXueUseRepeat
	JueXueUseNoExist
	JueXueActivateNotice
)

var (
	juexueLangMap = map[LangCode]string{
		JueXueRepeatActive:        "该绝学已激活,无需激活",
		JueXueNotActiveNotUpgrade: "未激活的绝学,无法升级",
		JueXueReacheFullUpgrade:   "绝学已达最高级",
		JueXueNotActiveNotInsight: "未激活的绝学,无法顿悟",
		JueXueInsightNotReach:     "条件未达成, 无法顿悟",
		JueXueUseNotActive:        "该绝学还未激活,无法使用",
		JueXueUseRepeat:           "绝学正在使用",
		JueXueUseNoExist:          "绝学使用不存在",
		JueXueActivateNotice:      "恭喜%s成功激活绝学%s，打架PK无往不利",
	}
)

func init() {
	mergeLang(juexueLangMap)
}
