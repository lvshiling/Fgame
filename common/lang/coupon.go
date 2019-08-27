package lang

const (
	CouponExchangeFailed LangCode = CouponBase + iota
)

var couponLangMap = map[LangCode]string{
	CouponExchangeFailed: "兑换失败",
}

func init() {
	mergeLang(couponLangMap)
}
