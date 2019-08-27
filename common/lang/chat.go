package lang

const (
	ChatLevelOrVipLow LangCode = ChatBase + iota
	ChatForbidden
)

var (
	chatLangMap = map[LangCode]string{
		ChatLevelOrVipLow: "等级达到%s级或VIP达到%s级方可在此频道进行发言",
		ChatForbidden:     "每天2:00~8:00无法使用聊天功能哦亲~",
	}
)

func init() {
	mergeLang(chatLangMap)
}
