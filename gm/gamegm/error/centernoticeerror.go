package error

const (
	ErrorCodeCenterLoginNoticeEmpty = baseCenterLoginNoticeError + iota
	ErrorCodeCenterLoginNoticeExist
)

var (
	errorCenterLoginNoticerMap = map[ErrorCode]string{
		ErrorCodeCenterLoginNoticeEmpty: "平台名不能为空",
		ErrorCodeCenterLoginNoticeExist: "平台名已经存在",
	}
)

var (
	minCenterLoginNoticeError = ErrorCodeCenterLoginNoticeEmpty
	maxCenterLoginNoticeError = ErrorCodeCenterLoginNoticeExist
)

func init() {
	for i := minCenterLoginNoticeError; i <= maxCenterLoginNoticeError; i++ {
		errorMsg := errorCenterLoginNoticerMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
