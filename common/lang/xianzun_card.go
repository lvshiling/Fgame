package lang

const (
	XianZunCardAlreadyActivite LangCode = XianZunCardBase + iota
	XianZunCardTemplateNotExist
	XianZunCardNotActivite
	XianZunCardAlreadyReceive
)

var (
	xianZunCardLangMap = map[LangCode]string{
		XianZunCardAlreadyActivite:  "您的仙尊特权卡已经激活",
		XianZunCardTemplateNotExist: "模板不存在",
		XianZunCardNotActivite:      "仙尊特权卡未激活",
		XianZunCardAlreadyReceive:   "您的仙尊特权卡每日奖励已经领取过了",
	}
)

func init() {
	mergeLang(xianZunCardLangMap)
}
