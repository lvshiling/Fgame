package error

const (
	ErrorCodeCenterPlatformEmpty = baseCenterPlatformError + iota
	ErrorCodeCenterPlatformExist
	ErrorCodeCenterPlatformSdkExist
)

var (
	errorCenterPlatformrMap = map[ErrorCode]string{
		ErrorCodeCenterPlatformEmpty:    "平台名不能为空",
		ErrorCodeCenterPlatformExist:    "平台名已经存在",
		ErrorCodeCenterPlatformSdkExist: "Sdk重复",
	}
)

var (
	minCenterPlatformError = ErrorCodeCenterPlatformEmpty
	maxCenterPlatformError = ErrorCodeCenterPlatformSdkExist
)

func init() {
	for i := minCenterPlatformError; i <= maxCenterPlatformError; i++ {
		errorMsg := errorCenterPlatformrMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
