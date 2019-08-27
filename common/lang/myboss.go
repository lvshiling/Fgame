package lang

const (
	MyBossBaseVipNotEnough LangCode = MyBossBase + iota
	MyBossBaseTimesNotEnough
)

var myBossLangMap = map[LangCode]string{
	MyBossBaseVipNotEnough:   "您当前VIP等级小于%s，无法进入！",
	MyBossBaseTimesNotEnough: "您本日挑战该BOSS次数已达上限，无法继续进行挑战",
}

func init() {
	mergeLang(myBossLangMap)
}
