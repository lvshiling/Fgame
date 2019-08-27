package handler

import (
	clientcodec "fgame/fgame/client/codec"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
)

func init() {
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SYSTEM_MESSAGE_TYPE), (*uipb.SCSystemMessage)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HEARTBEAT_TYPE), (*uipb.CSHeartBeat)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HEARTBEAT_TYPE), (*uipb.SCHeartBeat)(nil))
}
