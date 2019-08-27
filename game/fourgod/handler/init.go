package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOURGOD_GET_TYPE), (*uipb.SCFourGodGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOURGOD_OPEN_BOX_TYPE), (*uipb.CSFourGodOpenBox)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOURGOD_KEYNUM_CHANGE_TYPE), (*uipb.SCFourGodKeyNumChange)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOURGOD_BIO_BROADCAST_TYPE), (*uipb.SCFourGodBioBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOURGOD_TOTAL_TYPE), (*uipb.SCFourGodTotal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOURGOD_USE_MASKED_TYPE), (*uipb.CSFourGodUseMasked)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOURGOD_USE_MASKED_TYPE), (*uipb.SCFourGodUseMasked)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOURGOD_OPEN_BOX_TYPE), (*uipb.SCFourGodOpenBox)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOURGOD_OPEN_BOX_STOP_TYPE), (*uipb.SCFourGodOpenBoxStop)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOURGOD_OPEN_BOX_FINISH_TYPE), (*uipb.SCFourGodOpenBoxFinish)(nil))
}
