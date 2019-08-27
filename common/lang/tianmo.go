package lang

const (
	TianMoEatDanReachedLimit = TianMoBase + iota
	TianMoEatDanReachedFull
	TianMoAdanvacedReachedLimit
	TianMoAdvanceToLow
	TianMoAdvanceToHigh
	TianMoAdvanceNotEqual
	TianMoActivateNotEnoughCharge
)

var (
	tianMoLangMap = map[LangCode]string{
		TianMoEatDanReachedLimit:      "天魔体丹食丹等级已达最大,请进阶后再试",
		TianMoEatDanReachedFull:       "天魔体丹食丹等级满级",
		TianMoAdanvacedReachedLimit:   "天魔体已达最高阶",
		TianMoAdvanceToLow:            "您天魔体系统的阶别不足，无法使用物品",
		TianMoAdvanceToHigh:           "您天魔体系统的阶别过高，无法使用物品",
		TianMoAdvanceNotEqual:         "您天魔体系统的阶别不符，无法使用物品",
		TianMoActivateNotEnoughCharge: "充值数不满足激活条件",
	}
)

func init() {
	mergeLang(tianMoLangMap)
}
