package msg

import (
	mglog "fgame/fgame/gm/gamegm/gm/mglog/metadata"
)

func init() {
	rst := make([]*mglog.MsgItemInfo, 0)
	var item *mglog.MsgItemInfo

	item = &mglog.MsgItemInfo{
		Label:      "当前金额",
		DataColumn: "curMoney",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "变化前金额",
		DataColumn: "beforeMoney",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "变化的金额",
		DataColumn: "changed",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "升级原因编号",
		DataColumn: "reason",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	item = &mglog.MsgItemInfo{
		Label:      "升级原因",
		DataColumn: "reasonText",
		ShowType:   "normal",
	}
	rst = append(rst, item)

	mglog.RegisterMsgItemInfo("player_feedbackfee", "兑换", mglog.MsgTypePlayerLog, rst)
}
