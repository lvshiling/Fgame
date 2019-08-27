package error

const (
	ErrorCodeOpenApiEmpty = baseApiError + iota
	ErrorCodeOpenApiParam
	ErrorCodeOpenApiParamMiss
	ErrorCodeOpenApiSign
	ErrorCodeOpenApiParamBindError
	ErrorCodeOpenApiUserNotExists
	ErrorCodeOpenApiNotPrivilege
	ErrorCodeOpenApiParamNotEmpty
	ErrorCodeOpenApiSupportPoolGoldNumNotEqual
	ErrorCodeOpenApiSupportUserNotExists
	ErrorCodeOpenApiUnknown
)

var (
	errorOpenApiMap = map[ErrorCode]string{
		ErrorCodeOpenApiEmpty:                      "用户ID不能为空",
		ErrorCodeOpenApiParam:                      "参数异常",
		ErrorCodeOpenApiParamMiss:                  "缺失参数",
		ErrorCodeOpenApiSign:                       "sign签名异常",
		ErrorCodeOpenApiParamBindError:             "参数异常,json的key和value都必须为字符串",
		ErrorCodeOpenApiUserNotExists:              "userId不对或者缺失",
		ErrorCodeOpenApiNotPrivilege:               "非扶持玩家",
		ErrorCodeOpenApiParamNotEmpty:              "参数不能为空",
		ErrorCodeOpenApiSupportPoolGoldNumNotEqual: "金额与数量的个数必须一致",
		ErrorCodeOpenApiSupportUserNotExists:       "用户不存在",
		ErrorCodeOpenApiUnknown:                    "未知错误",
	}
)

var (
	minOpenApiError = ErrorCodeOpenApiEmpty
	maxOpenApiError = ErrorCodeOpenApiUnknown
)

func init() {
	for i := minOpenApiError; i <= maxOpenApiError; i++ {
		errorMsg := errorOpenApiMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
