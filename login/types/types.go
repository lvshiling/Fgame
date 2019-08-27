package types

type RealNameState int32

const (
	//未实名
	RealNameStateNone RealNameState = iota
	RealNameStateUnder18
	RealNameStateUp18
)
