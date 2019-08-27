package error

const (
	ErrorCodeChatSetEmpty = baseCenterChatSetError + iota
	ErrorCodeChatSetExist
	ErrorCodeChatSetRemote
)

var (
	errorChatSetrMap = map[ErrorCode]string{
		ErrorCodeChatSetEmpty:  "服务器不能为空",
		ErrorCodeChatSetExist:  "服务器配置已经存在",
		ErrorCodeChatSetRemote: "远程配置异常",
	}
)

var (
	minChatSetError = ErrorCodeChatSetEmpty
	maxChatSetError = ErrorCodeChatSetRemote
)

func init() {
	for i := minChatSetError; i <= maxChatSetError; i++ {
		errorMsg := errorChatSetrMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
