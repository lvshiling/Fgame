package handler

import "fgame/fgame/common/codec"
import uipb "fgame/fgame/common/codec/pb/ui"

import accountcodec "fgame/fgame/account/codec"

func init() {
	accountcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SYSTEM_MESSAGE_TYPE), (*uipb.SCSystemMessage)(nil))
	accountcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EXCEPTION_TYPE), (*uipb.SCException)(nil))
}
