package lang

const (
	DragonEatReachLimit LangCode = GemBase + iota
	DragonActiveNoEnough
	DragonAdvancedRewTitle
	DragonAdvancedRewContent
)

var (
	dragonLangMap = map[LangCode]string{
		DragonEatReachLimit:      "当前食用已达上限",
		DragonActiveNoEnough:     "条件不足",
		DragonAdvancedRewTitle:   "神龙现世升阶奖励",
		DragonAdvancedRewContent: "神龙现世升阶奖励内容",
	}
)

func init() {
	mergeLang(dragonLangMap)
}
