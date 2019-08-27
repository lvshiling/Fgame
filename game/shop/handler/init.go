package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHOP_LIMIT_TYPE), (*uipb.CSShopLimit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHOP_LIMIT_TYPE), (*uipb.SCShopLimit)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHOP_BUY_TYPE), (*uipb.CSShopBuy)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHOP_BUY_TYPE), (*uipb.SCShopBuy)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHOP_AUTO_BUY_LIST_TYPE), (*uipb.SCShopAutoBuyList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHOP_STOP_AUTO_BUY_TYPE), (*uipb.SCShopStopAutoBuy)(nil))
}
