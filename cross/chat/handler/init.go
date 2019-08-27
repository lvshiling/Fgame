package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHAT_SEND_TYPE), (*uipb.CSChatSend)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHAT_SEND_TYPE), (*uipb.SCChatSend)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHAT_RECV_TYPE), (*uipb.SCChatRecv)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHAT_LIST_NOTICE_TYPE), (*uipb.SCChatListNotice)(nil))
}
