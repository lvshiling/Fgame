package lang

const (
	XinFaRepeatActive LangCode = XinFaBase + iota
	XinFaNotActiveNotUpgrade
	XinFaReacheFullUpgrade
	XinFaActivateNotice
)

var (
	xinfaLangMap = map[LangCode]string{
		XinFaRepeatActive:        "该心法已激活,无需激活",
		XinFaNotActiveNotUpgrade: "未激活的心法,无法升级",
		XinFaReacheFullUpgrade:   "心法已达最高级",
		XinFaActivateNotice:      "恭喜%s成功激活心法%s，打架PK无往不利",
	}
)

func init() {
	mergeLang(xinfaLangMap)
}
