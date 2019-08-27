package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DIANXING_GET_TYPE), (*uipb.CSDianxingGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DIANXING_GET_TYPE), (*uipb.SCDianxingGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DIANXING_XINGCHEN_VARY_TYPE), (*uipb.SCDianxingXingchenVary)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DIANXING_ADVANCED_TYPE), (*uipb.CSDianxingAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DIANXING_ADVANCED_TYPE), (*uipb.SCDianxingAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DIANXING_JIEFENG_ADVANCED_TYPE), (*uipb.CSDianxingJiefengAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DIANXING_JIEFENG_ADVANCED_TYPE), (*uipb.SCDianxingJiefengAdvanced)(nil))
}
