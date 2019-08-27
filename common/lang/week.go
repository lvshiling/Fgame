package lang

const (
	WeekHadBuyWeek LangCode = WeekBase + iota
	WeekNotBuyWeek
	WeekHadReceiveRewards
)

var (
	weekLangMap = map[LangCode]string{
		WeekHadBuyWeek:        "已经购买周卡",
		WeekNotBuyWeek:        "未购买周卡",
		WeekHadReceiveRewards: "已领取今日奖励",
	}
)

func init() {
	mergeLang(weekLangMap)
}
