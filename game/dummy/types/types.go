package types

type DummyType int32

const (
	DummyTypeSurname    DummyType = iota + 1 //姓氏
	DummyTypeFemaleName                      //女性名字
	DummyTypeMaleName                        //男性名字
)

func (t DummyType) Valid() bool {
	switch t {
	case DummyTypeFemaleName, DummyTypeMaleName, DummyTypeSurname:
		return true
	default:
		return false
	}
}
