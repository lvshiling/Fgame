package trade

import (
	"fgame/fgame/common/lang"
	gamecommon "fgame/fgame/game/common/common"
)

var (
	//个人上架商品已达上限
	errorTradeItemPersonalNumLimit = gamecommon.CodeError(lang.TradeItemPersonalNumLimit)
	//上架商品已达上限
	errorTradeItemTotalNumLimit = gamecommon.CodeError(lang.TradeItemTotalNumLimit)
	//商品已经不存在
	errorTradeItemNoExist = gamecommon.CodeError(lang.TradeItemNoExist)
	//上架商品已经不存在
	errorTradeUploadItemNoExist = gamecommon.CodeError(lang.TradeUploadItemNoExist)
	//上架商品还没上架
	errorTradeUploadItemNoUpload = gamecommon.CodeError(lang.TradeUploadItemNoUpload)
	//已经被其他人下单了
	errorTradeItemAlreadyOrderOther = gamecommon.CodeError(lang.TradeItemAlreadyOrderOther)
	//已经被自己下单了
	errorTradeItemAlreadyOrderSelf = gamecommon.CodeError(lang.TradeItemAlreadyOrderSelf)
)
