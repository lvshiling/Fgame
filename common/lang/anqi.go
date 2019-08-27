package lang

const (
	AnqiEatDanReachedLimit = AnqiBase + iota
	AnqiEatDanReachedFull
	AnqiAdanvacedReachedLimit
	AnqiAdvanceToLow
	AnqiAdvanceToHigh
	AnqiAdvanceNotEqual
	AnqiAdvancedNotice
)

var (
	anqiLangMap = map[LangCode]string{
		AnqiEatDanReachedLimit:    "暗器丹食丹等级已达最大,请进阶后再试",
		AnqiEatDanReachedFull:     "暗器丹食丹等级满级",
		AnqiAdanvacedReachedLimit: "暗器已达最高阶",
		AnqiAdvanceToLow:          "您暗器系统的阶别不足，无法使用物品",
		AnqiAdvanceToHigh:         "您暗器系统的阶别过高，无法使用物品",
		AnqiAdvanceNotEqual:       "您暗器系统的阶别不符，无法使用物品",
		AnqiAdvancedNotice:        "淬毒致命，%s成功将暗器提升至%s，战力飙升%s，攻击时附加的毒药更为强大了",
	}
)

func init() {
	mergeLang(anqiLangMap)
}
