package lang

const (
	LivenessOpenBoxNoEnough LangCode = LivenessBase + iota
	LivenessRewardTitle
	LivenessRewardContent
)

var (
	livenessLangMap = map[LangCode]string{
		LivenessOpenBoxNoEnough: "条件不足或已领取完",
		LivenessRewardTitle:     "活跃度奖励",
		LivenessRewardContent:   "您成功完成了活跃度目标,但奖励未进行领取,系统自动补发给您,敬请领取",
	}
)

func init() {
	mergeLang(livenessLangMap)
}
