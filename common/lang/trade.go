package lang

const (
	TradeUploadRefundTitle LangCode = TradeBase + iota
	TradeUploadRefundContent
	TradeUploadWithdrawTitle
	TradeUploadWithdrawContent
	TradeUploadWithdrawBySystemContent
	TradeItemRefundTitle
	TradeItemRefundContent
	TradeItemTitle
	TradeItemContent
	TradeSellTitle
	TradeSellContent
	TradeItemPriceTooLow
	TradeItemPriceTooHigh
	TradeItemNoExist
	TradeItemPersonalNumLimit
	TradeItemTotalNumLimit
	TradeUploadItemNoExist
	TradeUploadItemNoUpload
	TradeItemAlreadyOrderOther
	TradeItemAlreadyOrderSelf
	TradeServiceClose
	TradeServiceNoPrivilege
)

var (
	tradeLangMap = map[LangCode]string{
		TradeUploadRefundTitle:             "上架失败",
		TradeUploadRefundContent:           "亲爱的玩家，交易行上架数量已达上限，上架物品返还给您，请稍后再试",
		TradeUploadWithdrawTitle:           "商品下架",
		TradeUploadWithdrawContent:         "商品下架",
		TradeUploadWithdrawBySystemContent: "您的商品由于无人问津，系统已自动将您的商品下架，请尽快收取！如需继续寄售，请重新前往交易市场上架！",
		TradeItemRefundTitle:               "交易市场返还",
		TradeItemRefundContent:             "由于网络波动原因，您在交易市场的购买操作没有成功，现将已消耗的元宝的返还给您，请查收！",
		TradeItemTitle:                     "购买成功",
		TradeItemContent:                   "购买成功",
		TradeSellTitle:                     "商品出售",
		TradeSellContent:                   "%s在交易市场上购买了您寄售的%s，物品数量%s，花费：%s（您当前VIP等级：%s，手续费：%s%%，需要支付手续费：%s），以下为您获得的收益，敬请查收",
		TradeItemPriceTooLow:               "价格太低",
		TradeItemPriceTooHigh:              "价格太高",
		TradeItemNoExist:                   "商品已经不存在",
		TradeItemPersonalNumLimit:          "当前市场寄售物品过多，无法继续上架物品，请稍后再来",
		TradeItemTotalNumLimit:             "当前市场寄售物品过多，无法继续上架物品，请稍后再来",
		TradeUploadItemNoExist:             "上架商品已经不存在",
		TradeUploadItemNoUpload:            "上架商品还没上架",
		TradeItemAlreadyOrderOther:         "商品已经被别人下单了",
		TradeItemAlreadyOrderSelf:          "商品已经被自己下单了",
		TradeServiceClose:                  "交易行维护中",
		TradeServiceNoPrivilege:            "交易行没有权限",
	}
)

func init() {
	mergeLang(tradeLangMap)
}
