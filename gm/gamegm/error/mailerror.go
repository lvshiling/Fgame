package error

const (
	ErrorCodeEmailEmpty = baseEmailError + iota
	ErrorCodeEmailChecked
)

var (
	errorEmailrMap = map[ErrorCode]string{
		ErrorCodeEmailEmpty:   "空邮件",
		ErrorCodeEmailChecked: "邮件已被审核",
	}
)

var (
	minEmailError = ErrorCodeEmailEmpty
	maxEmailError = ErrorCodeEmailChecked
)

func init() {
	for i := minEmailError; i <= maxEmailError; i++ {
		errorMsg := errorEmailrMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
