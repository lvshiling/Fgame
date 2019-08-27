package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XINFA_GET_TYPE), (*uipb.SCXinFaGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XINFA_ACTIVE_TYPE), (*uipb.CSXinFaActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XINFA_ACTIVE_TYPE), (*uipb.SCXinFaActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XINFA_UPGRADE_TYPE), (*uipb.CSXinFaUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XINFA_UPGRADE_TYPE), (*uipb.SCXinFaUpgrade)(nil))

}
