package lang

const (
	AccountNoExist LangCode = AccountBase + iota
	AccountLoginAtSameTime
	AccountLoginFailed
	AccountLoginRepeat
	AccountRegisterClose
	AccountArgumentInvalid
	AccountOrIpForbid
)

var accountLangMap = map[LangCode]string{
	AccountNoExist:         "玩家登录信息失效，请重试！",
	AccountLoginAtSameTime: "您的账号在其他设备登录，若非您本人操作，建议修改登录密码！",
	AccountLoginFailed:     "登陆失败",
	AccountLoginRepeat:     "重复登陆",
	AccountRegisterClose:   "服务器已满,建议前往新服更火爆!",
	AccountArgumentInvalid: "账户参数错误",
	AccountOrIpForbid:      "账户或ip被封",
}

func init() {
	mergeLang(accountLangMap)
}
