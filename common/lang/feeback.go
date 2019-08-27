package lang

const (
	FeebackExchange LangCode = FeebackBase + iota
	FeebackMoneyNoEnoughOrLimit
	FeebackMoneyMoneyWrong
)

var (
	feebackLangMap = map[LangCode]string{
		FeebackExchange:             "兑换中",
		FeebackMoneyNoEnoughOrLimit: "余额不足或已经超过限制",
		FeebackMoneyMoneyWrong:      "兑换金额有误",
	}
)

func init() {
	mergeLang(feebackLangMap)
}
