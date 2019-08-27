package types

type XianTaoType int32

const (
	XianTaoTypeQianNian XianTaoType = iota
	XianTaoTypeBaiNian
)

func (t XianTaoType) Valid() bool {
	switch t {
	case XianTaoTypeQianNian,
		XianTaoTypeBaiNian:
		return true
	}
	return false
}

const (
	XianTaoTypeMin = XianTaoTypeQianNian
	XianTaoTypeMax = XianTaoTypeBaiNian
)
