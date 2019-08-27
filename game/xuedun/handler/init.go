package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	initCodec()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XUEDUN_GET_TYPE), (*uipb.CSXueDunGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUEDUN_GET_TYPE), (*uipb.SCXueDunGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XUEDUN_UPGRADE_TYPE), (*uipb.CSXueDunUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUEDUN_UPGRADE_TYPE), (*uipb.SCXueDunUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XUEDUN_PEIYANGE_TYPE), (*uipb.CSXueDunPeiYang)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUEDUN_PEIYANGE_TYPE), (*uipb.SCXueDunPeiYang)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XUEDUN_BLOOD_CHANGED_TYPE), (*uipb.SCXueDunBloodChanged)(nil))
}
