package handler

import "fgame/fgame/common/codec"
import crosscodec "fgame/fgame/cross/codec"
import uipb "fgame/fgame/common/codec/pb/ui"

func init() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ERROR_TYPE), (*uipb.SCError)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EXCEPTION_TYPE), (*uipb.SCException)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SYSTEM_MESSAGE_TYPE), (*uipb.SCSystemMessage)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GET_TIME_TYPE), (*uipb.CSGetTime)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GET_TIME_TYPE), (*uipb.SCGetTime)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HEARTBEAT_TYPE), (*uipb.CSHeartBeat)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HEARTBEAT_TYPE), (*uipb.SCHeartBeat)(nil))
}
