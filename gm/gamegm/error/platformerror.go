package error

const (
	ErrorCodePlatformEmpty = basePlatformError + iota
	ErrorCodePlatformExistSdk
	ErrorCodePlatformExist
)

var (
	errorPlatformrMap = map[ErrorCode]string{
		ErrorCodePlatformEmpty:    "平添名称渠道等不能为空",
		ErrorCodePlatformExistSdk: "sdk已存在",
		ErrorCodePlatformExist:    "平台名已经存在",
	}
)

var (
	minPlatformError = ErrorCodePlatformEmpty
	maxPlatformError = ErrorCodePlatformExist
)

func init() {
	for i := minPlatformError; i <= maxPlatformError; i++ {
		errorMsg := errorPlatformrMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
