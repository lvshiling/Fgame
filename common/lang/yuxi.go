package lang

const (
	YuXiWinTitle LangCode = YuXiBase + iota
	YuXiWinContent
	YuXiWinNotWinerMember
	YuXiWinHadReceiveDayRew
	YuXiWinNoticeContent
	YuXiWinHadNotDayRew
	YuXiOwnerCanNotTransfer
)

var (
	yuXiLangMap = map[LangCode]string{
		YuXiWinTitle:            "玉玺之战",
		YuXiWinContent:          "玉玺之战获胜奖励",
		YuXiWinNotWinerMember:   "不是玉玺之战获胜仙盟成员",
		YuXiWinHadReceiveDayRew: "玉玺之战每日奖励已领取",
		YuXiWinHadNotDayRew:     "玉玺之战没有可领取的每日奖励",
		YuXiWinNoticeContent:    "恭喜%s仙盟获得本次玉玺争夺战的胜利，仙盟成员将享受四大特权！",
		YuXiOwnerCanNotTransfer: "玉玺持有者不可接受邀请",
	}
)

func init() {
	mergeLang(yuXiLangMap)
}
