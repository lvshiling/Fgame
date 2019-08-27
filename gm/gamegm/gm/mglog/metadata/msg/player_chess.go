package msg

import (
	mglog "fgame/fgame/gm/gamegm/gm/mglog/metadata"
)

func init() {
	rst := make([]*mglog.MsgItemInfo, 0)
	var item *mglog.MsgItemInfo

	item = &mglog.MsgItemInfo{
		Label:      "抽奖次数",
		DataColumn: "attendTimes",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "原因编号",
		DataColumn: "reason",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "日志原因",
		DataColumn: "reasonText",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	mglog.RegisterMsgItemInfo("player_chess", "苍龙棋局", mglog.MsgTypePlayerLog, rst)
}
