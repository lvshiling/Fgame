package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/chat/chat"
	chattypes "fgame/fgame/game/chat/types"
)

func BuildSCChatRecv(sendId int64, channel chattypes.ChannelType, recvId int64, msgType chattypes.MsgType, content []byte) *uipb.SCChatRecv {
	chatRecv := &uipb.SCChatRecv{}
	chatRecv.SendId = &sendId
	channelInt := int32(channel)
	chatRecv.Channel = &channelInt
	chatRecv.RecvId = &recvId
	msgTypeInt := int32(msgType)
	chatRecv.MsgType = &msgTypeInt
	chatRecv.Content = content

	return chatRecv
}

func BuildSCChatRecvWithCliArgs(sendId int64, sendName string, channel chattypes.ChannelType, recvId int64, msgType chattypes.MsgType, content []byte, args string) *uipb.SCChatRecv {
	chatRecv := &uipb.SCChatRecv{}
	chatRecv.SendId = &sendId
	channelInt := int32(channel)
	chatRecv.Channel = &channelInt
	chatRecv.RecvId = &recvId
	msgTypeInt := int32(msgType)
	chatRecv.MsgType = &msgTypeInt
	chatRecv.Content = content
	chatRecv.Args = &args
	chatRecv.SendName = &sendName
	return chatRecv
}

func BuildSCChatSend(channel chattypes.ChannelType, recvId int64, msgType chattypes.MsgType, content []byte, args string) *uipb.SCChatSend {
	chatSend := &uipb.SCChatSend{}

	channelInt := int32(channel)
	chatSend.Channel = &channelInt
	chatSend.RecvId = &recvId
	msgTypeInt := int32(msgType)
	chatSend.MsgType = &msgTypeInt
	chatSend.Content = content
	chatSend.Args = &args
	return chatSend
}

func BuildSCChatSendWithChatCount(channel chattypes.ChannelType, recvId int64, msgType chattypes.MsgType, content []byte, args string, chatCount int32) *uipb.SCChatSend {
	chatSend := &uipb.SCChatSend{}

	channelInt := int32(channel)
	chatSend.Channel = &channelInt
	chatSend.RecvId = &recvId
	msgTypeInt := int32(msgType)
	chatSend.MsgType = &msgTypeInt
	chatSend.Content = content
	chatSend.Args = &args
	chatSend.ChatCount = &chatCount
	return chatSend
}

func BuildSCWorldChatListNotice(worldChatList, systemChatList []*chat.ChatData) *uipb.SCChatListNotice {
	scWorldChatListNotice := &uipb.SCChatListNotice{}

	for _, data := range worldChatList {
		scWorldChatListNotice.WorldList = append(scWorldChatListNotice.WorldList, buildChatData(data))
	}
	for _, data := range systemChatList {
		scWorldChatListNotice.SystemList = append(scWorldChatListNotice.SystemList, buildChatData(data))
	}

	return scWorldChatListNotice
}

func buildChatData(data *chat.ChatData) *uipb.ChatData {
	chatData := &uipb.ChatData{}
	msgTypeInt := int32(data.GetMsgType())
	sendId := data.GetSendId()
	content := data.GetContent()
	sendTime := data.GetSendTime()
	args := data.GetArgs()

	chatData.SendId = &sendId
	chatData.MsgType = &msgTypeInt
	chatData.Content = content
	chatData.SendTime = &sendTime
	chatData.Args = &args

	return chatData
}
