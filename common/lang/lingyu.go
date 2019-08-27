package lang

const (
	LingyuUnrealDanReachedLimit LangCode = LingyuBase + iota
	LingyuUnrealCondNotReached
	LingyuUnrealDanReachedFull
	LingyuAdanvacedReachedLimit
	LingyuUnrealNoExist
	LingyuAdvanceToLow
	LingyuAdvanceToHigh
	LingyuAdvanceNotEqual
	LingyuAdvancedNotice
	LingyuSkinUpstarNoActive
	LingyuSkinReacheFullStar
)

var (
	lingyuLangMap = map[LangCode]string{
		LingyuUnrealDanReachedLimit: "幻化丹食丹等级已达最大,请进阶后再试",
		LingyuUnrealDanReachedFull:  "幻化丹食丹等级满级",
		LingyuUnrealCondNotReached:  "还有幻化条件未达成，无法解锁幻化",
		LingyuAdanvacedReachedLimit: "已达最高阶",
		LingyuUnrealNoExist:         "当前没有幻化",
		LingyuAdvanceToLow:          "您领域系统的阶别不足，无法使用物品",
		LingyuAdvanceToHigh:         "您领域系统的阶别过高，无法使用物品",
		LingyuAdvanceNotEqual:       "您领域系统的阶别不符，无法使用物品",
		LingyuAdvancedNotice:        "战神发威，%s成功将领域提升至%s，战力飙升%s，加入仙盟战神堂可分享领域战力！",
		LingyuSkinUpstarNoActive:    "未激活的领域皮肤,无法升星",
		LingyuSkinReacheFullStar:    "领域皮肤已满星",
	}
)

func init() {
	mergeLang(lingyuLangMap)
}
