package types

type MarryPriceType int32

const (
	MarryPriceTypeNone MarryPriceType = iota
	MarryPriceTypeNormal
	MarryPriceTypeCheap
	MarryPriceTypeExp
)

type ZhiZunType int32

const (
	ZhiZunTypeNormal ZhiZunType = iota + 1 //正常版本
	ZhiZunTypeExp                          //贵价版本
)

func (t ZhiZunType) Valid() bool {
	switch t {
	case ZhiZunTypeExp,
		ZhiZunTypeNormal:
		return true
	}
	return false
}
