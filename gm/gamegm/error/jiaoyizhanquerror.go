package error

const (
	ErrorCodeJiaoYiZhanQuEmpty = baseJiaoYiZhanQuError + iota
	ErrorCodeJiaoYiZhanQuExists
)

var (
	errorJiaoYiZhanQuMap = map[ErrorCode]string{
		ErrorCodeJiaoYiZhanQuEmpty:  "空",
		ErrorCodeJiaoYiZhanQuExists: "已经存在",
	}
)

var (
	minJiaoYiError = ErrorCodeJiaoYiZhanQuEmpty
	maxJiaoYiError = ErrorCodeJiaoYiZhanQuExists
)

func init() {
	for i := minJiaoYiError; i <= maxJiaoYiError; i++ {
		errorMsg := errorJiaoYiZhanQuMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
