package lang

const (
	XianTiEatUnDanReachedLimit LangCode = XianTiBase + iota
	XianTiEatUnDanReachedFull
	XianTiUnrealCondNotReached
	XianTiAdanvacedReachedLimit
	XianTiUnrealNoExist
	XianTiOtherIsNoActived
	XianTiAdvanceToLow
	XianTiAdvanceToHigh
	XianTiAdvanceNotEqual
	XianTiSkinUpstarNoActive
	XianTiSkinReacheFullStar
	XianTiAdvancedNotice
	XianTiUnrealActivateNotice
)

var (
	xianTiLangMap = map[LangCode]string{
		XianTiEatUnDanReachedLimit:  "幻化丹食丹等级已达最大,请进阶后再试",
		XianTiEatUnDanReachedFull:   "幻化丹食丹等级满级",
		XianTiUnrealCondNotReached:  "还有条件未达成，无法解锁幻化",
		XianTiAdanvacedReachedLimit: "仙体已达最高阶",
		XianTiUnrealNoExist:         "当前没有幻化",
		XianTiOtherIsNoActived:      "该仙体还未激活",
		XianTiAdvanceToLow:          "您仙体系统的阶别不足，无法使用物品",
		XianTiAdvanceToHigh:         "您仙体系统的阶别过高，无法使用物品",
		XianTiAdvanceNotEqual:       "您仙体系统的阶别不符，无法使用物品",
		XianTiSkinUpstarNoActive:    "未激活的仙体皮肤,无法升星",
		XianTiSkinReacheFullStar:    "仙体皮肤已满星",
		XianTiAdvancedNotice:        "恭喜%s成功将仙体提升至%s，战力飙升%s",
		XianTiUnrealActivateNotice:  "恭喜%s成功将仙体幻化为%s，战力飙升%s",
	}
)

func init() {
	mergeLang(xianTiLangMap)
}
