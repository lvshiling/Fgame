package lang

const (
	DianXingAdanvacedReachedLimit = DianXingBase + iota
	DianXingAdvanceNotXingChen
	DianXingAdvancedNotice
	DianXingJieFengAdanvacedReachedLimit
)

var (
	dianXingLangMap = map[LangCode]string{
		DianXingAdanvacedReachedLimit:        "点星已达最高级",
		DianXingAdvanceNotXingChen:           "您的星尘数量不足，无法进行操作",
		DianXingAdvancedNotice:               "恭喜%s成功点星%s星谱，星芒万丈，属性飙升！%s",
		DianXingJieFengAdanvacedReachedLimit: "点星解封已达最高级",
	}
)

func init() {
	mergeLang(dianXingLangMap)
}
