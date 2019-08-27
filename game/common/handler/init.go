package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/processor"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ERROR_TYPE), (*uipb.SCError)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EXCEPTION_TYPE), (*uipb.SCException)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SYSTEM_MESSAGE_TYPE), (*uipb.SCSystemMessage)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GET_TIME_TYPE), (*uipb.CSGetTime)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GET_TIME_TYPE), (*uipb.SCGetTime)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HEARTBEAT_TYPE), (*uipb.CSHeartBeat)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HEARTBEAT_TYPE), (*uipb.SCHeartBeat)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_OPEN_SERVER_TIME_TYPE), (*uipb.CSOpenServerTime)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_OPEN_SERVER_TIME_TYPE), (*uipb.SCOpenServerTime)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MERGE_SERVER_TIME_TYPE), (*uipb.CSMergeServerTime)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MERGE_SERVER_TIME_TYPE), (*uipb.SCMergeServerTime)(nil))

}

func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_GET_TIME_TYPE))
}
