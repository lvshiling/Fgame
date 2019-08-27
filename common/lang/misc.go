package lang

const (
	MiscRealNameAuthAlready LangCode = MiscBase + iota
	MiscRealNameAuthCodeWrong
	MiscRealNamePhoneInvalid
	MiscRealNameAuthInfoWrong
	MiscExitKaSiCd
)

var (
	miscLangMap = map[LangCode]string{
		MiscRealNameAuthAlready:   "已经实名认证过了",
		MiscRealNameAuthCodeWrong: "手机验证码错误",
		MiscRealNamePhoneInvalid:  "手机号码无效",
		MiscRealNameAuthInfoWrong: "身份证和名字匹配不上",
		MiscExitKaSiCd:            "操作太频繁,请%s秒后再试",
	}
)

func init() {
	mergeLang(miscLangMap)
}
