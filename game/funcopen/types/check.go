package types

type FuncOpenCheckType int32

const (
	//自动开启
	FuncOpenCheckTypeAuto FuncOpenCheckType = iota
	//手动开启
	FuncOpenCheckTypeManual
)

func (spt FuncOpenCheckType) Valid() bool {
	switch spt {
	case FuncOpenCheckTypeAuto,
		FuncOpenCheckTypeManual:
		return true
	}
	return false
}
