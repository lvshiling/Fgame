/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*ChatContent)(nil))
}

/*聊天内容*/
type ChatContent struct {
	PlayerLogMsg `bson:",inline"`

	//频道(0:世界,1:帮派,2:队伍,3:系统,4:私聊)
	Channel int32 `json:"channel"`

	//接收者id
	RecvId int64 `json:"recvId"`

	//接收者名字
	RecvName string `json:"recvName"`

	//消息类型(0:文本,1:表情,2:语音)
	MsgType int32 `json:"msgType"`

	//内容
	Content []byte `json:"content"`

	//文本内容
	Text string `json:"text"`
}

func (c *ChatContent) LogName() string {
	return "chat_content"
}
