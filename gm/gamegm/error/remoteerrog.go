package error

const (
	ErrorCodeRemoteEmpty = baseRemoteError + iota
	ErrorCodeDefaultRemoteUser
	ErrorCodeDefaultRemoteCenter
	ErrorCodeDefaultRemoteNotice
)

var (
	errorRemoteMap = map[ErrorCode]string{
		ErrorCodeRemoteEmpty:         "remote服务为空",
		ErrorCodeDefaultRemoteUser:   "中心用户服务异常",
		ErrorCodeDefaultRemoteCenter: "中心服务异常",
		ErrorCodeDefaultRemoteNotice: "中心通知服务异常",
	}
)

var (
	minRemoteError = ErrorCodeRemoteEmpty
	maxRemoteError = ErrorCodeDefaultRemoteNotice
)

func init() {
	for i := minRemoteError; i <= maxRemoteError; i++ {
		errorMsg := errorRemoteMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
