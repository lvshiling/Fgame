package handler

import "fgame/fgame/common/codec"
import uipb "fgame/fgame/common/codec/pb/ui"

import accountcodec "fgame/fgame/account/codec"

func init() {
	accountcodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ACCOUNT_LOGIN_TYPE), (*uipb.CSAccountLogin)(nil))
	accountcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ACCOUNT_LOGIN_TYPE), (*uipb.SCAccountLogin)(nil))
}
