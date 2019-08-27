package lang

const (
	VipHadBuyGift LangCode = VipBase + iota
	VipLevelToLow
	VipHadReceiveFreeGift
)

var (
	vipLangMap = map[LangCode]string{
		VipHadBuyGift:         "该vip礼包不能重复购买",
		VipLevelToLow:         "vip等级不足",
		VipHadReceiveFreeGift: "vip免费礼包已领取",
	}
)

func init() {
	mergeLang(vipLangMap)
}
