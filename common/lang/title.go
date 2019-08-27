package lang

const (
	TitleRepeatActive LangCode = TitleBase + iota
	TitleRepeatWear
	TitleNotHas
	TitleActivateNotice
	TitleNotValid
	TitleTemplateNotExist
	TitleStarLevelAlreadyTop
	TitleNotEver
)

var (
	titleLangMap = map[LangCode]string{
		TitleRepeatActive:        "称号已激活,无需激活",
		TitleRepeatWear:          "称号已穿戴,无需再次穿戴",
		TitleNotHas:              "没有该称号,请先获取",
		TitleActivateNotice:      "%s成功激活%s称号，战力暴涨%s",
		TitleNotValid:            "不是限时称号，无法重复使用",
		TitleTemplateNotExist:    "模板不存在",
		TitleStarLevelAlreadyTop: "升星等级已经满级",
		TitleNotEver:             "不是永久称号，无法升星",
	}
)

func init() {
	mergeLang(titleLangMap)
}
