package error

const (
	ErrorCodeChannelEmpty = baseChannelError + iota
	ErrorCodeChannelExist
)

var (
	errorChannelrMap = map[ErrorCode]string{
		ErrorCodeChannelEmpty: "渠道名不能为空",
		ErrorCodeChannelExist: "渠道名已经存在",
	}
)

var (
	minChannelError = ErrorCodeChannelEmpty
	maxChannelError = ErrorCodeChannelExist
)

func init() {
	for i := minChannelError; i <= maxChannelError; i++ {
		errorMsg := errorChannelrMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
