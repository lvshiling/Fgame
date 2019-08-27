package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUNC_OPEN_LIST_TYPE), (*uipb.SCFuncOpenList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUNC_OPEN_UPDATE_LIST_TYPE), (*uipb.SCFuncOpenUpdateList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FUNC_OPEN_MANUAL_ACTIVE_TYPE), (*uipb.CSFuncOpenManualActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUNC_OPEN_MANUAL_ACTIVE_TYPE), (*uipb.SCFuncOpenManualActive)(nil))
}
