package lang

const (
	LongGongEmailTitle LangCode = LongGongBase + iota
	LongGongEmailContent
	LongGongCollectCountNoEnough
)

var (
	longGongLangMap = map[LangCode]string{
		LongGongEmailTitle:           "龙宫探宝",
		LongGongEmailContent:         "您在龙宫探宝活动中，对黑龙至尊的伤害排第%d名，以下是您的排名奖励，请查收！",
		LongGongCollectCountNoEnough: "采集次数上限",
	}
)

func init() {
	mergeLang(longGongLangMap)
}
