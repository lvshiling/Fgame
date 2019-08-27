package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_REAL_NAME_AUTH_TYPE), (*uipb.CSRealNameAuth)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REAL_NAME_AUTH_TYPE), (*uipb.SCRealNameAuth)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GET_IDENTIFY_CODE_TYPE), (*uipb.CSGetIdentifyCode)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GET_IDENTIFY_CODE_TYPE), (*uipb.SCGetIdentifyCode)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EXIT_KASI_TYPE), (*uipb.CSExitKaSi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EXIT_KASI_TYPE), (*uipb.SCExitKaSi)(nil))

}
