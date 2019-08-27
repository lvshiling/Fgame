package types

//日志接口
type LogReason interface {
	//类型
	Reason() int32
	//描述
	String() string
}
