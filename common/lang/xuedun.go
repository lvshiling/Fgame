package lang

const (
	XueDunUpgradeReachFull LangCode = XueDunBase + iota
	XueDunUpgradeNoBlood
	XueDunPeiYangReachFull
	XueDunEatCulDanReachedLimit
	XueDunUpgradeNotice
)

var (
	xuedunLangMap = map[LangCode]string{
		XueDunUpgradeReachFull:      "血玉系统已达满阶",
		XueDunUpgradeNoBlood:        "血炼值不足,无法升阶",
		XueDunPeiYangReachFull:      "血玉培养最高级",
		XueDunEatCulDanReachedLimit: "血玉培养等级已达最大,请升阶后再试",
		XueDunUpgradeNotice:         "血玉护体,恭喜%s成功将血玉系统提升至%s,血量大涨,生命值提升%s",
	}
)

func init() {
	mergeLang(xuedunLangMap)
}
