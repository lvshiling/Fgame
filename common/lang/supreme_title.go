package lang

const (
	SupremeTitleRepeatActive LangCode = SupremeTitleBase + iota
	SupremeTitleRepeatWear
	SupremeTitleNotHas
	SupremeTitleActivateNotice
)

var (
	supremeTitleLangMap = map[LangCode]string{
		SupremeTitleRepeatActive:   "至尊称号已激活,无需激活",
		SupremeTitleRepeatWear:     "至尊称号已穿戴,无需再次穿戴",
		SupremeTitleNotHas:         "没有该至尊称号,请先获取",
		SupremeTitleActivateNotice: "%s成功激活%s至尊称号，战力暴涨%s",
	}
)

func init() {
	mergeLang(supremeTitleLangMap)
}
