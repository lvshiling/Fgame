package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EMPEROR_GET_TYPE), (*uipb.CSEmperorGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EMPEROR_GET_TYPE), (*uipb.SCEmperorGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EMPEROR_WORSHIP_TYPE), (*uipb.CSEmperorWorship)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EMPEROR_WORSHIP_TYPE), (*uipb.SCEmperorWorship)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EMPEROR_STORAGE_TYPE), (*uipb.CSEmperorStorageGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EMPEROR_STORAGE_TYPE), (*uipb.SCEmperorStorageGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EMPEROR_ROB_TYPE), (*uipb.CSEmperorRob)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EMPEROR_ROB_TYPE), (*uipb.SCEmperorRob)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EMPEROR_RECORD_TYPE), (*uipb.CSEmperorRecords)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EMPEROR_RECORD_TYPE), (*uipb.SCEmperorRecords)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EMPEROR_ROBBED_TYPE), (*uipb.SCEmperorRobbed)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EMPEROR_OPEN_BOX_TYPE), (*uipb.CSEmperorOpenBox)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EMPEROR_OPEN_BOX_TYPE), (*uipb.SCEmperorOPenBox)(nil))
}
