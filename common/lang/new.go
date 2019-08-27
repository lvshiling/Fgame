package lang

const (
	NewTitle LangCode = NewBase + iota
	NewContent
)

var newLangMap = map[LangCode]string{
	NewTitle:   "新手礼包",
	NewContent: "欢庆开服，全民福利！这是系统小姐姐亲手为您准备的一份小礼物哦（价值388元宝），愿您修仙路上一帆风顺",
}

func init() {
	mergeLang(newLangMap)
}
