package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GUA_JI_TYPE), (*uipb.CSGuaJi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUA_JI_TYPE), (*uipb.SCGuaJi)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CURRENT_GUA_JI_TYPE), (*uipb.SCCurrentGuaJi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_STOP_GUA_JI_TYPE), (*uipb.CSStopGuaJi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_STOP_GUA_JI_TYPE), (*uipb.SCStopGuaJi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUA_JI_POS_TYPE), (*uipb.SCGuaJiPos)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GUA_JI_ADVANCE_LIST_TYPE), (*uipb.CSGuaJiAdvanceList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUA_JI_ADVANCE_LIST_TYPE), (*uipb.SCGuaJiAdvanceList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUA_JI_ADVANCE_UPDATE_LIST_TYPE), (*uipb.SCGuaJiAdvanceUpdateList)(nil))
}
