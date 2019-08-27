package msg

import (
	mglog "fgame/fgame/gm/gamegm/gm/mglog/metadata"
)

func init() {
	rst := make([]*mglog.MsgItemInfo, 0)
	var item *mglog.MsgItemInfo

	item = &mglog.MsgItemInfo{
		Label:      "之前声威值",
		DataColumn: "beforeShengWeiZhi",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "当前声威值",
		DataColumn: "curShengWeiZhi",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "进阶原因编号",
		DataColumn: "reason",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "进阶原因",
		DataColumn: "reasonText",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	mglog.RegisterMsgItemInfo("player_jieyi_shengweizhi", "结义声威值改变", mglog.MsgTypePlayerLog, rst)
}
