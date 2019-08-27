package lang

const (
	ShiHunFanEatDanReachedLimit = ShiHunFanBase + iota
	ShiHunFanEatDanReachedFull
	ShiHunFanAdanvacedReachedLimit
	ShiHunFanAdvancedNotCharge
	ShiHunFanAdvanceNotEqual
)

var (
	shiHunFanLangMap = map[LangCode]string{
		ShiHunFanEatDanReachedLimit:    "噬魂幡培养等级已达最大,请进阶后再试",
		ShiHunFanEatDanReachedFull:     "噬魂幡培养等级满级",
		ShiHunFanAdanvacedReachedLimit: "噬魂幡已达最高阶",
		ShiHunFanAdvancedNotCharge:     "您噬魂幡系统的充值数不足",
		ShiHunFanAdvanceNotEqual:       "您噬魂幡系统的阶别不符，无法使用物品",
	}
)

func init() {
	mergeLang(shiHunFanLangMap)
}
