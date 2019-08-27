package handler

import (
	clientcodec "fgame/fgame/client/codec"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
)

func init() {
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ACCOUNT_LOGIN_TYPE), (*uipb.CSAccountLogin)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ACCOUNT_LOGIN_TYPE), (*uipb.SCAccountLogin)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LOGIN_TYPE), (*uipb.CSLogin)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LOGIN_TYPE), (*uipb.SCLogin)(nil))

	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEST_LOGIN_TYPE), (*uipb.CSTestLogin)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEST_LOGIN_TYPE), (*uipb.SCTestLogin)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ENETER_SELECT_JOB_TYPE), (*uipb.SCEnterSelectJob)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SELECT_JOB_TYPE), (*uipb.CSSelectJob)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SELECT_JOB_TYPE), (*uipb.SCSelectJob)(nil))

}
