package error

//错误编码管理

type ErrorCode int

type Error interface {
	error
	Code() ErrorCode
}

type singleError struct {
	code     ErrorCode
	errorMsg string
}

func (m *singleError) Code() ErrorCode {
	return m.code
}

func (m *singleError) Error() string {
	return m.errorMsg
}

func NewError(p_code ErrorCode, errorMsg string) Error {
	rst := &singleError{
		code:     p_code,
		errorMsg: errorMsg,
	}
	return rst
}

const (
	baseServerError            ErrorCode = 1000  //服务器处理异常编码
	baseUserError              ErrorCode = 2000  //用户处理异常编码
	baseChannelError           ErrorCode = 3000  //渠道处理异常编码
	basePlatformError          ErrorCode = 4000  //平台处理异常编码
	baseGmServerSupportPool    ErrorCode = 5000  //扶植池错误编码
	baseCenterPlatformError    ErrorCode = 20000 //中心平台处理异常编码
	baseCenterServerError      ErrorCode = 21000 //中心服务器处理异常编码
	baseCenterChatSetError     ErrorCode = 22000 //中心服务器处理异常编码
	baseCenterLoginNoticeError ErrorCode = 23000 //中心服务登陆日志异常
	baseEmailError             ErrorCode = 24000 //邮件异常
	baseRemoteError            ErrorCode = 99000 //remote异常
	baseApiError               ErrorCode = 98000 //API
	baseJiaoYiZhanQuError      ErrorCode = 25000
)

var (
	errorMap = make(map[ErrorCode]error)
)

func addError(code ErrorCode, err error) {
	errorMap[code] = err
}

func GetError(code ErrorCode) error {
	err, ok := errorMap[code]
	if !ok {
		return nil
	}
	return err
}
