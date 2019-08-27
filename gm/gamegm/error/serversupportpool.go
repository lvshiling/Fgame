package error

const (
	ErrorCodeServerSupportPoolEmpty = baseGmServerSupportPool + iota
	ErrorCodeServerSupportPoolExists
	ErrorCodeServerSupportPoolNotExists
	ErrorCodeServerSupportPoolNotEnought
	ErrorCodePlatformSupportPoolSetExists
)

var (
	errorServerSupportPoolMap = map[ErrorCode]string{
		ErrorCodeServerSupportPoolEmpty:       "",
		ErrorCodeServerSupportPoolExists:      "服务器已经存在",
		ErrorCodeServerSupportPoolNotExists:   "扶持池不存在",
		ErrorCodeServerSupportPoolNotEnought:  "扶持池元宝不足",
		ErrorCodePlatformSupportPoolSetExists: "平台扶持配置已存在",
	}
)

var (
	minServerSupportPoolError = ErrorCodeServerSupportPoolEmpty
	maxServerSupportPoolError = ErrorCodeServerSupportPoolNotEnought
)

func init() {
	for i := minServerSupportPoolError; i <= maxServerSupportPoolError; i++ {
		errorMsg := errorServerSupportPoolMap[i]
		addError(i, NewError(i, errorMsg))
	}
}
