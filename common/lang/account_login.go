package lang

const (
	AccountLoginPlatformInvalid LangCode = AccountLoginBase + iota
	AccountLoginPlatformVerifyFailed
	AccountLoginMessageCanNotHandle
)

var accountLoginLangMap = map[LangCode]string{
	AccountLoginPlatformInvalid:      "平台无效",
	AccountLoginPlatformVerifyFailed: "平台验证失败",
	AccountLoginMessageCanNotHandle:  "消息异常",
}

func init() {
	mergeLang(accountLoginLangMap)
}
