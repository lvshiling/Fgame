package lang

const (
	ShopBuyNumInvalid LangCode = ShopBase + iota
	ShoBuyReacheLimit
	ShopBuyNotItem
	ShopAdvancedAutoBuyItemFail
	ShopUpstarAutoBuyItemFail
	ShopChessAutoBuyItemFail
	ShopOpenLightAutoBuyItemFail
	ShopFireworksAutoBuyItemFail
	ShopFlowerAutoBuyItemFail
	ShopReliveAutoBuyItemFail
	ShopXueChiAutoBuyItemFail
	ShopMingGeAutoBuyItemFail
	ShopMingLiAutoBuyItemFail
	ShopZhenFaShengJiAutoBuyItemFail
	ShopZhenQiXianHuoAutoBuyItemFail
	ShopBabyLearnAutoBuyItemFail
)

var (
	shopLangMap = map[LangCode]string{
		ShopBuyNumInvalid:                "购买数量大于最大购买数量",
		ShoBuyReacheLimit:                "购买次数，已达每日限购数量",
		ShopBuyNotItem:                   "商铺没有该道具,无法自动购买",
		ShopAdvancedAutoBuyItemFail:      "购买物品失败,自动进阶已停止",
		ShopUpstarAutoBuyItemFail:        "购买物品失败,自动升星已停止",
		ShopOpenLightAutoBuyItemFail:     "购买物品失败,自动开光已停止",
		ShopChessAutoBuyItemFail:         "购买物品失败,棋局破解失败",
		ShopMingGeAutoBuyItemFail:        "购买物品失败,自动祭炼失败",
		ShopMingLiAutoBuyItemFail:        "购买物品失败,洗练自动购买已停止",
		ShopZhenFaShengJiAutoBuyItemFail: "购买物品失败,阵法自动升级已停止",
		ShopZhenQiXianHuoAutoBuyItemFail: "购买物品失败,阵法仙火自动升级停止",
		ShopFireworksAutoBuyItemFail:     "购买烟花物品失败",
		ShopFlowerAutoBuyItemFail:        "购买鲜花物品失败",
		ShopReliveAutoBuyItemFail:        "购买复活丹物品失败",
		ShopXueChiAutoBuyItemFail:        "购买血瓶物品失败",
		ShopBabyLearnAutoBuyItemFail:        "购买宝宝读书物品失败",
	}
)

func init() {
	mergeLang(shopLangMap)
}
