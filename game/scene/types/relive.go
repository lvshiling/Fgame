package types


type ReliveType int32
const (
	//原地复活
	ReliveTypeImmediate ReliveType = iota
	//回城复活
	ReliveTypeBack
	//回进入点
	ReliveTypeEnterPoint
	//回复活点
	ReliveTypeRelivePoint
)

func (t ReliveType) Mask() int32 {
	return 1 << uint(t)
}

func (t ReliveType) Valid() bool {
	switch t {
	case ReliveTypeImmediate,
		ReliveTypeBack,
		ReliveTypeEnterPoint,
		ReliveTypeRelivePoint:
		return true
	}
	return false
}
