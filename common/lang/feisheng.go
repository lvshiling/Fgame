package lang

const (
	FeiShengFullRate LangCode = FeiShengBase + iota
	FeiShengGongDeNotEnough
	FeiShengHadLeftQn
	FeiShengSanGongNotice
	FeiShengLevelToLower
	FeiShengLimit
)

var (
	feiLangMap = map[LangCode]string{
		FeiShengFullRate:        "成功率已达100%，无需再次服用金丹",
		FeiShengGongDeNotEnough: "功德不足，请先获取功德",
		FeiShengHadLeftQn:       "还未进行加点，无法保存",
		FeiShengLevelToLower:    "飞升等级不足",

		FeiShengSanGongNotice: "%s在%s进行散功！将自身修为散去，分发给同一地图其他玩家，并获得大量功德奖励！%s",
		FeiShengLimit:         "已达本日可接受散功经验次数上限，请明日再来",
	}
)

func init() {
	mergeLang(feiLangMap)
}
