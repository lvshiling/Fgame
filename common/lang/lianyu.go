package lang

const (
	LianYuCheckFuncOpen LangCode = LianYuBase + iota
	LianYuInLineUpEnterFuBen
	LianYuCancleLineUpNoExist
	LianYuBossRefreshNotice
	LianYuBossDeadNotice
)

var (
	lianYuLangMap = map[LangCode]string{
		LianYuCheckFuncOpen:       "您当前未开启无间炼狱功能,无法进入活动",
		LianYuInLineUpEnterFuBen:  "您当前正在无间炼狱排队中,无法进入副本",
		LianYuCancleLineUpNoExist: "您当前未在排队中",
		LianYuBossRefreshNotice:   "%s已经复活!大家赶紧前往击杀BOSS,获取杀气吧",
		LianYuBossDeadNotice:      "%s已被%s击杀",
	}
)

func init() {
	mergeLang(lianYuLangMap)
}
