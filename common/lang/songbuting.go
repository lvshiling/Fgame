package lang

const (
	SongBuTingNoReceive LangCode = SongBuTingBase + iota
	SongBuTingReceiveNumLimit
	SongBuTingTitle
	SongBuTingContent
	SongBuTingNotice
)

var (
	songBuTingLangMap = map[LangCode]string{
		SongBuTingNoReceive:       "单笔充值%d元宝,才能享受享尊贵特权",
		SongBuTingReceiveNumLimit: "今日福利已领完,请明日再来",
		SongBuTingTitle:           "%d元宝送不停",
		SongBuTingContent:         "恭喜您在开服活动中单笔充值金额达到%d元宝,每天均可领取丰厚奖励,以下为您的奖励,敬请查收",
		SongBuTingNotice:          "恭喜%s在%s中单笔充值%s元宝，成功获取至尊特权,一辈子获得绑定元宝返还,每天可领取%d绑元!",
	}
)

func init() {
	mergeLang(songBuTingLangMap)
}
