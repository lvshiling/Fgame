package lang

const (
	ChargeOrderFailed LangCode = ChargeBase + iota
	ChargePrivilegeFailed
	ChargeFirstChargeNoDone
)

var (
	chargeLangMap = map[LangCode]string{
		ChargeOrderFailed:       "下单失败",
		ChargePrivilegeFailed:   "后台扶持失败",
		ChargeFirstChargeNoDone: "首充活动还没结束",
	}
)

func init() {
	mergeLang(chargeLangMap)
}
