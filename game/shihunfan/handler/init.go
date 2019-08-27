package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHIHUNFAN_GET_TYPE), (*uipb.CSShihunfanGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHIHUNFAN_GET_TYPE), (*uipb.SCShihunfanGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHIHUNFAN_CHARGE_VARY_TYPE), (*uipb.SCShihunfanChargeVary)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHIHUNFAN_ADVANCED_TYPE), (*uipb.CSShihunfanAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHIHUNFAN_ADVANCED_TYPE), (*uipb.SCShihunfanAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHIHUNFAN_DAN_ADVANCED_TYPE), (*uipb.CSShihunfanDanAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHIHUNFAN_DAN_ADVANCED_TYPE), (*uipb.SCShihunfanDanAdvanced)(nil))
}
