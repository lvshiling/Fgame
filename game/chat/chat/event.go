package chat

import chattypes "fgame/fgame/game/chat/types"

type ChatEventData struct {
	channel  chattypes.ChannelType
	recvId   int64
	recvName string
	msgType  chattypes.MsgType
	content  []byte
}

func (d *ChatEventData) GetChannel() chattypes.ChannelType {
	return d.channel
}

func (d *ChatEventData) GetRecvId() int64 {
	return d.recvId
}

func (d *ChatEventData) GetRecvName() string {
	return d.recvName
}

func (d *ChatEventData) GetMsgType() chattypes.MsgType {
	return d.msgType
}

func (d *ChatEventData) GetContent() []byte {
	return d.content
}

func CreateChatEventData(channel chattypes.ChannelType, recvId int64, recvName string, msgType chattypes.MsgType, content []byte) *ChatEventData {
	d := &ChatEventData{}
	d.channel = channel
	d.recvId = recvId
	d.recvName = recvName
	d.msgType = msgType
	d.content = content
	return d
}
