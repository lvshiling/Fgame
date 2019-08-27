package lang

const (
	HuiYuanHadBuyHuiYuan = HuiYuanBase + iota
	HuiYuanNotHuiYuan
	HuiYuanHadReceiveRewards
	HuiYuanBuyNotice
)

var (
	huiyuanLangMap = map[LangCode]string{
		HuiYuanHadBuyHuiYuan:     "已购买会员",
		HuiYuanNotHuiYuan:        "不是会员",
		HuiYuanHadReceiveRewards: "已领取奖励",
		HuiYuanBuyNotice:         "身份尊贵，%s成功激活永久至尊会员，每天免费领取%s",
	}
)

func init() {
	mergeLang(huiyuanLangMap)
}
