package lang

const (
	FashionRepeatActive LangCode = FashionBase + iota
	FashionRepeatWear
	FashionNotHas
	FashionBornNoUnload
	FashionNotActiveNotUpstar
	FashionReacheFullStar
	FashionActivateNotice
	FashionTrialHadActivate
	FashionTrialCardUseIsExist
)

var (
	fashionLangMap = map[LangCode]string{
		FashionRepeatActive:        "时装已激活,无需激活",
		FashionRepeatWear:          "时装已穿戴,无需再次穿戴",
		FashionNotHas:              "没有该时装,请先获取",
		FashionBornNoUnload:        "出生自带时装不能卸下",
		FashionNotActiveNotUpstar:  "未激活的时装,无法升星",
		FashionReacheFullStar:      "时装已满星",
		FashionActivateNotice:      "%s成功激活酷炫时装%s，战力暴涨%s",
		FashionTrialHadActivate:    "您当前已激活永久时装，无法使用体验卡",
		FashionTrialCardUseIsExist: "已获得时装试用,期间无法重复获取",
	}
)

func init() {
	mergeLang(fashionLangMap)
}
