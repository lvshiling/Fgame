package lang

const (
	MoonloveNotSinglePlayer LangCode = MoonloveBase + iota
	MoonloveNotInScene
	MoonloveIsCoupleUseItem
	MoonloveCollectTimesNotEnough
	MoonloveFindFail
	MoonloveSameSex
)

var (
	moonloveLangMap = map[LangCode]string{
		MoonloveNotSinglePlayer:       "该玩家已和他人赏月中！",
		MoonloveNotInScene:            "玩家不在月下情缘场景",
		MoonloveIsCoupleUseItem:       "当前处于双人赏月状态，无法使用该物品",
		MoonloveCollectTimesNotEnough: "当前剩余采集次数不足，无法采集",
		MoonloveFindFail:              "该玩家当前已离开该场景，无法双人赏月！",
		MoonloveSameSex:               "仅能选择异性玩家进行双人赏月！",
	}
)

func init() {
	mergeLang(moonloveLangMap)
}
