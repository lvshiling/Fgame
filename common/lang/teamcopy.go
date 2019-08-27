package lang

const (
	TeamCopyTitle LangCode = TeamCopyBase + iota
	TeamCopyContent
)

var (
	teamCopyLangMap = map[LangCode]string{
		TeamCopyTitle:   "组队副本",
		TeamCopyContent: "由于您背包空间不足,无法领取奖励,系统以邮件形式进行发送,敬请查收",
	}
)

func init() {
	mergeLang(teamCopyLangMap)
}
