package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOUND_RESOUCE_LIST_TYPE), (*uipb.CSFoundResouceList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUND_RESOUCE_LIST_TYPE), (*uipb.SCFoundResouceList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOUND_TYPE), (*uipb.CSFound)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUND_TYPE), (*uipb.SCFound)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOUND_BATCH_TYPE), (*uipb.CSFoundBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOUND_BATCH_TYPE), (*uipb.SCFoundBatch)(nil))
}
