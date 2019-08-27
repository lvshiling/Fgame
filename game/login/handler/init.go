package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LOGIN_TYPE), (*uipb.CSLogin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LOGIN_TYPE), (*uipb.SCLogin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEST_LOGIN_TYPE), (*uipb.CSTestLogin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEST_LOGIN_TYPE), (*uipb.SCTestLogin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ENETER_SELECT_JOB_TYPE), (*uipb.SCEnterSelectJob)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SELECT_JOB_TYPE), (*uipb.CSSelectJob)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SELECT_JOB_TYPE), (*uipb.SCSelectJob)(nil))
}
