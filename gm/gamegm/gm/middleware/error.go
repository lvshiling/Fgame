package middleware

const (
	//无权限
	ErrorCodeNoPrivilege = 1000 + iota
)

var (
	errorMap = map[int]string{
		ErrorCodeNoPrivilege: "没有权限",
	}
)
