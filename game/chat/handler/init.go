package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHAT_SEND_TYPE), (*uipb.CSChatSend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHAT_SEND_TYPE), (*uipb.SCChatSend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHAT_RECV_TYPE), (*uipb.SCChatRecv)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHAT_LIST_NOTICE_TYPE), (*uipb.SCChatListNotice)(nil))
}
