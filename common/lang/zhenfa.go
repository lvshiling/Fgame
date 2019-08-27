package lang

const (
	ZhenFaActivated LangCode = ZhenFaBase + iota
	ZhenFaActivatePreCond
	ZhenFaShengJiNoActivate
	ZhenFaShengJiFullLevel
	ZhenQiAdvancedNoActivate
	ZhenQiAdvancedFullLevel
	ZhenQiXianHuoShengJiNoActivate
	ZhenFaXianHuoFullLevel
	ZhenFaZhenQiAdvancedNoEnoughLevel
)

var (
	zhenFaLangMap = map[LangCode]string{
		ZhenFaActivated:                   "该阵法已经激活过",
		ZhenFaActivatePreCond:             "需要%s阵法达到%s级",
		ZhenFaShengJiNoActivate:           "未激活阵法,无法升级",
		ZhenFaShengJiFullLevel:            "阵法已达最高级",
		ZhenQiAdvancedNoActivate:          "当前阵法未激活,无法升阶阵旗",
		ZhenQiAdvancedFullLevel:           "当前阶别已达最高",
		ZhenQiXianHuoShengJiNoActivate:    "当前阵法未激活,无法升级仙火",
		ZhenFaXianHuoFullLevel:            "阵旗仙火已达最高级",
		ZhenFaZhenQiAdvancedNoEnoughLevel: "需要%s阵法达到%s级",
	}
)

func init() {
	mergeLang(zhenFaLangMap)
}
