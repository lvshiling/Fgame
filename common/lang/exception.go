package lang

const (
	ExceptionPlayerStateError = ExceptionBase + iota
	ExceptionPlayerAuthTimeout
	ExceptionServerBusy
)

var (
	exceptionLangMap = map[LangCode]string{
		ExceptionPlayerStateError:  "玩家状态异常",
		ExceptionPlayerAuthTimeout: "玩家认证超时",
		ExceptionServerBusy:        "服务器繁忙",
	}
)

func init() {
	mergeLang(exceptionLangMap)
}
