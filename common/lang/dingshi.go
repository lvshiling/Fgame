package lang

const (
	DingShiBossBlood LangCode = DingShiBase + iota
)

var (
	dingShiLangMap = map[LangCode]string{
		DingShiBossBlood: "世界BOSS%s在血量仅剩%d%%，击杀它可获得海量珍稀道具，赶紧前往争夺！%s",
	}
)

func init() {
	mergeLang(dingShiLangMap)
}
