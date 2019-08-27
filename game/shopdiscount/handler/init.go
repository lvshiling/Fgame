package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHOP_DISCOUNT_GET_TYPE), (*uipb.CSShopDiscountGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHOP_DISCOUNT_GET_TYPE), (*uipb.SCShopDiscountGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHOP_DISCOUNT_NOTICE_TYPE), (*uipb.SCShopDiscountNotice)(nil))
}
