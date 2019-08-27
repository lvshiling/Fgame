package error

const (
	ErrorCodeCenterServerEmpty = baseCenterServerError + iota
	ErrorCodeCenterServerIdExist
	ErrorCodeCenterServerExist
	ErrorCodeCenterServerDBNotConnect
	ErrorCodeCenterServerRemoteNotConnect
)

var (
	errorCenterServerrMap = map[ErrorCode]string{
		ErrorCodeCenterServerEmpty:            "服务器名和中心平台不能为空",
		ErrorCodeCenterServerIdExist:          "服务器ID已经存在",
		ErrorCodeCenterServerExist:            "服务器名称名已经存在",
		ErrorCodeCenterServerDBNotConnect:     "配置成功，但DB无法连接",
		ErrorCodeCenterServerRemoteNotConnect: "配置成功，但Remote无法连接",
	}
)

var (
	minCenterServerError = ErrorCodeCenterServerEmpty
	maxCenterServerError = ErrorCodeCenterServerRemoteNotConnect
)

func init() {
	for i := minCenterServerError; i <= maxCenterServerError; i++ {
		errorMsg := errorCenterServerrMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
