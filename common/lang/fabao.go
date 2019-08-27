package lang

const (
	FaBaoUnrealDanReachedLimit LangCode = FaBaoBase + iota
	FaBaoUnrealDanReachedFull
	FaBaoUnrealCondNotReached
	FaBaoAdanvacedReachedLimit
	FaBaoUnrealNoExist
	FaBaoOtherIsNoActived
	FaBaoAdvanceToLow
	FaBaoAdvanceToHigh
	FaBaoAdvanceNotEqual
	FaBaoSkinUpstarNoActive
	FaBaoSkinReacheFullStar
	FaBaoAdvancedNotice
	FaBaoUnrealActivateNotice
	FaBaoUnrealActiveSystem
	FaBaoHadActivate
	FaBaoTongLingReachedFull
)

var (
	faBaoLangMap = map[LangCode]string{
		FaBaoUnrealDanReachedLimit: "幻化丹食丹等级已达最大,请进阶后再试",
		FaBaoUnrealDanReachedFull:  "幻化丹食丹等级满级",
		FaBaoUnrealCondNotReached:  "还有幻化条件未达成，无法解锁幻化",
		FaBaoAdanvacedReachedLimit: "已达最高阶",
		FaBaoUnrealNoExist:         "当前没有幻化",
		FaBaoOtherIsNoActived:      "该法宝还未激活",
		FaBaoAdvanceToLow:          "您法宝系统的阶别不足，无法使用物品",
		FaBaoAdvanceToHigh:         "您法宝系统的阶别过高，无法使用物品",
		FaBaoAdvanceNotEqual:       "您法宝系统的阶别不符，无法使用物品",
		FaBaoSkinUpstarNoActive:    "未激活的法宝皮肤,无法升星",
		FaBaoSkinReacheFullStar:    "法宝皮肤已满星",
		FaBaoAdvancedNotice:        "恭喜%s成功将法宝提升至%s，战力飙升%s",
		FaBaoUnrealActivateNotice:  "恭喜%s成功将法宝幻化为%s，战力飙升%s",
		FaBaoUnrealActiveSystem:    "请先激活法宝系统",
		FaBaoHadActivate:           "您当前已激活法宝系统",
		FaBaoTongLingReachedFull:   "通灵等级满级",
	}
)

func init() {
	mergeLang(faBaoLangMap)
}
