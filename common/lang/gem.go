package lang

const (
	GemMineFirstFlushActive LangCode = GemBase + iota
	GemNotReceiveStone
	GemGambleNotYuanShi
	GemMineActivateNotice
)

var (
	gemLangMap = map[LangCode]string{
		GemMineFirstFlushActive: "请先激活第二个矿工",
		GemNotReceiveStone:      "当前无原石可供领取,稍后再来",
		GemGambleNotYuanShi:     "原石不足,无法赌石",
		GemMineActivateNotice:   "家里有矿没人挖，%s狠狠心激活了一名矿工，原石产量直接翻倍",
	}
)

func init() {
	mergeLang(gemLangMap)
}
