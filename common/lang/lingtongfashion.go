package lang

const (
	LingTongFashionRepeateActivate LangCode = LingTongFashionBase + iota
	LingTongFashionUseNoActivate
	LingTongFashionBornNoUnload
	LingTongFashionUpstarNoActivate
	LingTongFashionReacheFullStar
	LingTongFashionTrialHadActivate
	LingTongFashionTrialCardUseIsExist
)

var (
	lingTongFashionLangMap = map[LangCode]string{
		LingTongFashionRepeateActivate:     "灵童时装重复激活",
		LingTongFashionUseNoActivate:       "未激活的灵童时装,无法使用",
		LingTongFashionBornNoUnload:        "出生自带时装不能卸下",
		LingTongFashionUpstarNoActivate:    "未激活的灵童时装,无法升星",
		LingTongFashionReacheFullStar:      "时装已满星",
		LingTongFashionTrialHadActivate:    "您当前已激活永久时装,无法使用体验卡",
		LingTongFashionTrialCardUseIsExist: "已获得时装试用,期间无法重复获取",
	}
)

func init() {
	mergeLang(lingTongFashionLangMap)
}
