package lang

const (
	ServerBusy LangCode = ServerBase + iota
	ServerNoOpen
)

var serverMap = map[LangCode]string{
	ServerBusy:   "服务器繁忙,请稍后重试",
	ServerNoOpen: "服务器将于%s对外开放,请稍后重试",
}

func init() {
	mergeLang(serverMap)
}
