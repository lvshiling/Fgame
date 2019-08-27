package lang

const (
	ShenfaUnrealDanReachedLimit LangCode = ShenfaBase + iota
	ShenfaUnrealCondNotReached
	ShenfaAdanvacedReachedLimit
	ShenfaUnrealNoExist
	ShenfaUnrealDanReachedFull
	ShenfaAdvanceToLow
	ShenfaAdvanceToHigh
	ShenfaAdvanceNotEqual
	ShenfaAdvancedNotice
	ShenfaSkinUpstarNoActive
	ShenfaSkinReacheFullStar
)

var (
	shenfaLangMap = map[LangCode]string{
		ShenfaUnrealDanReachedLimit: "食用幻化丹药数量已达最大,请进阶后再试",
		ShenfaUnrealDanReachedFull:  "幻化丹食丹等级满级",
		ShenfaUnrealCondNotReached:  "还有幻化条件未达成，无法解锁幻化",
		ShenfaAdanvacedReachedLimit: "已达最高阶",
		ShenfaUnrealNoExist:         "当前没有幻化",
		ShenfaAdvanceToLow:          "您身法系统的阶别不足，无法使用物品",
		ShenfaAdvanceToHigh:         "您身法系统的阶别过高，无法使用物品",
		ShenfaAdvanceNotEqual:       "您身法系统的阶别不符，无法使用物品",
		ShenfaAdvancedNotice:        "凌波微步，%s成功将身法提升至%s，战力飙升%s，战斗中几率闪避一切攻击",
		ShenfaSkinUpstarNoActive:    "未激活的身法皮肤,无法升星",
		ShenfaSkinReacheFullStar:    "身法皮肤已满星",
	}
)

func init() {
	mergeLang(shenfaLangMap)
}
