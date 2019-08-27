package lang

const (
	OneArenaGoGrab LangCode = OneArenaBase + iota
	OneArenaRobSameLevel
	OneArenaRobLowLevel
	OneArenaCoolTime
	OneArenaKunNoExist
	OneArenaRobSucess
	OneArenaOutputTitle
	OneArenaOutputContent
	OneArenaMergeServerTitle
	OneArenaMergeServerContent
)

var oneArenaLangMap = map[LangCode]string{
	OneArenaGoGrab:             "不能越级抢夺",
	OneArenaRobSameLevel:       "不能同级抢夺",
	OneArenaRobLowLevel:        "不能抢夺低级灵池",
	OneArenaCoolTime:           "您当前处于抢夺冷却时间中,请等冷却时间结束后再来!",
	OneArenaKunNoExist:         "当前没有鲲可以出售",
	OneArenaRobSucess:          "恭喜%s占领%s",
	OneArenaOutputTitle:        "灵池产出奖励",
	OneArenaOutputContent:      "灵池产出奖励内容",
	OneArenaMergeServerTitle:   "合服灵池争夺重置",
	OneArenaMergeServerContent: "您的服务器与其他服务器进行了合服,灵池争夺玩法进行重置,您失去了原先的灵池位置,您原来在仙灵店铺的产出已经帮您保留,请前往查收",
}

func init() {
	mergeLang(oneArenaLangMap)
}
