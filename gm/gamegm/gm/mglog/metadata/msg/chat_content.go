package msg

import (
	mglog "fgame/fgame/gm/gamegm/gm/mglog/metadata"
)

func init() {
	rst := make([]*mglog.MsgItemInfo, 0)
	var item *mglog.MsgItemInfo

	item = &mglog.MsgItemInfo{
		Label:      "频道",
		DataColumn: "channel",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "接收者id",
		DataColumn: "recvId",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "接收者名字",
		DataColumn: "recvName",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "消息类型",
		DataColumn: "msgType",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "内容",
		DataColumn: "content",
		ShowType:   "byte",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "文本内容",
		DataColumn: "text",
		ShowType:   "string",
	}
	rst = append(rst, item)

	mglog.RegisterMsgItemInfo("chat_content", "聊天内容", mglog.MsgTypePlayerLog, rst)
}
