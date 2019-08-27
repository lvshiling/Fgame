package lang

const (
	SystemCompensateMailTitle = SystemCompensateBase + iota
	SystemCompensateMailContent
)

var (
	systemLangMap = map[LangCode]string{
		SystemCompensateMailTitle:   "更新补偿",
		SystemCompensateMailContent: "尊敬的用户，由于当前版本%s开放至15阶，造成玩家战力损失，本次更新后将为损失战力的玩家发放补偿奖励。给您造成的不便我们深表歉意，我们将继续努力为大家带来更好的游戏体验。",
	}
)

func init() {
	mergeLang(systemLangMap)
}
