package msg

import (
	mglog "fgame/fgame/gm/gamegm/gm/mglog/metadata"
)

func init() {
	rst := make([]*mglog.MsgItemInfo, 0)
	var item *mglog.MsgItemInfo

	item = &mglog.MsgItemInfo{
		Label:      "变化元宝数",
		DataColumn: "changedNum",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "变化前的元宝数",
		DataColumn: "beforeGold",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "当前的元宝数",
		DataColumn: "curGold",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "变更原因编号",
		DataColumn: "reason",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "变更原因",
		DataColumn: "reasonText",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	mglog.RegisterMsgItemInfo("player_new_glod_changed", "新元宝变化", mglog.MsgTypePlayerLog, rst)
}
