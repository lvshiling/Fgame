package lang

const (
	ShangguzhilingLingShouUnLock LangCode = ShangguzhilingBase + iota
	ShangguzhilingLingShouFullLevel
	ShangguzhilingUseItemIdWrong
	ShangguzhilingUseItemCountNotEnough
	ShangguzhilingLingWenUnLock
	ShangguzhilingLingWenFullLevel
	ShangguzhilingLingShouFullRank
	ShangguzhilingLingLianUnLock
	ShangguzhilingReceiveCDNotEnough
)

var (
	shangguzhilingLangMap = map[LangCode]string{
		ShangguzhilingLingShouUnLock:        "上古之灵灵兽未解锁",
		ShangguzhilingLingShouFullLevel:     "上古之灵灵兽满级",
		ShangguzhilingUseItemIdWrong:        "上古之灵选择使用的物品ID错误",
		ShangguzhilingUseItemCountNotEnough: "上古之灵选择使用的物品数量不足",
		ShangguzhilingLingWenUnLock:         "上古之灵灵纹未解锁",
		ShangguzhilingLingWenFullLevel:      "上古之灵灵纹满级",
		ShangguzhilingLingShouFullRank:      "上古之灵灵兽已经满阶",
		ShangguzhilingLingLianUnLock:        "上古之灵灵炼部位未解锁，灵兽等级不足",
		ShangguzhilingReceiveCDNotEnough:    "奖励领取冷却中, 不能领取",
	}
)

func init() {
	mergeLang(shangguzhilingLangMap)
}
