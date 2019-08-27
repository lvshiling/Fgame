package lang

const (
	GodSiegeCheckFuncOpen LangCode = GodSiegeBase + iota
	GodSiegeInLineUpEnterFuBen
	GodSiegeCancleLineUpNoExist
	GodSiegeBossRefreshNotice
	GodSiegeBossDeadNotice
)

var (
	godSiegeLangMap = map[LangCode]string{
		GodSiegeCheckFuncOpen:       "您当前未开启神兽攻城功能,无法进入活动",
		GodSiegeInLineUpEnterFuBen:  "您当前正在神兽攻城排队中,无法进入副本",
		GodSiegeCancleLineUpNoExist: "您当前未在排队中",
		GodSiegeBossRefreshNotice:   "%s已经复活！大家赶紧前往击杀BOSS，获取%s系统属性材料",
		GodSiegeBossDeadNotice:      "%s已被%s击杀",
	}
)

func init() {
	mergeLang(godSiegeLangMap)
}
