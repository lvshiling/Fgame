package lang

const (
	EquipBaoKuNotGetRewards LangCode = EquipBaoKuBase + iota
	EquipBaoKuShopBuyNumInvalid
	EquipBaoKuShopBuyReacheLimit
	EquipBaoKuShopBuyNotItem
	EquipBaoKuShopBuyJiFenNotEnough
)

var (
	equipBaoKuLangMap = map[LangCode]string{
		EquipBaoKuNotGetRewards:         "运气不好，再来一次",
		EquipBaoKuShopBuyNumInvalid:     "兑换数量大于最大兑换数量",
		EquipBaoKuShopBuyReacheLimit:    "兑换次数，已达每日限购数量",
		EquipBaoKuShopBuyNotItem:        "商铺没有该道具,无法自动兑换",
		EquipBaoKuShopBuyJiFenNotEnough: "积分不足，无法完成兑换",
	}
)

func init() {
	mergeLang(equipBaoKuLangMap)
}
