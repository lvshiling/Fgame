package lang

const (
	ActivityNotAtTime LangCode = ActivityBase + iota
	ActivityNoTimes
	ActivityAlreadyClose
	ActivityStartNotice
	ActivityEndNotice
	ActivityTickRewEmailTitle
	ActivityTickRewEmailContent
	ActivityHadFinish
)

var (
	activityLangMap = map[LangCode]string{
		//------------------活动-------------------
		ActivityNotAtTime:           "当前活动未开启",
		ActivityNoTimes:             "当前活动次数不足",
		ActivityAlreadyClose:        "当前活动已关闭",
		ActivityStartNotice:         "还有%d分钟，%s活动就要开始了，请大家提前做好准备",
		ActivityEndNotice:           "还有%d分钟，%s活动就要结束了，请大家提前做好准备",
		ActivityTickRewEmailTitle:   "活动定时奖励",
		ActivityTickRewEmailContent: "背包空间不足!",
		ActivityHadFinish:           "本次活动已经结束!",
	}
)

func init() {
	mergeLang(activityLangMap)
}
