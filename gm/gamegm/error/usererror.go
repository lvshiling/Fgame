package error

const (
	ErrorCodeUserParamEmpty = baseUserError + iota
	ErrorCodeUserExist
	ErrorCodeMissChannel
	ErrorCodeMissPlatform
)

var (
	errorUserMap = map[ErrorCode]string{
		ErrorCodeUserParamEmpty: "用户名密码及权限等参数不能为空",
		ErrorCodeUserExist:      "用户名已经存在",
		ErrorCodeMissChannel:    "渠道缺失",
		ErrorCodeMissPlatform:   "平台缺失",
	}
)

var (
	minUserError = ErrorCodeUserParamEmpty
	maxUserError = ErrorCodeMissPlatform
)

func init() {
	for i := minUserError; i <= maxUserError; i++ {
		errorMsg := errorUserMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
