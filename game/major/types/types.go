package types

//副本类型
type MajorType int32

const (
	MajorTypeShuangXiu MajorType = iota //双修副本
	MajorTypeFuQi                       //夫妻副本
)

func (t MajorType) Valid() bool {
	switch t {
	case MajorTypeShuangXiu,
		MajorTypeFuQi:
		return true
	default:
		return false
	}
}

var (
	majorTypeMap = map[MajorType]string{
		MajorTypeShuangXiu: "双修副本",
		MajorTypeFuQi:      "夫妻副本",
	}
)

func (spt MajorType) String() string {
	return majorTypeMap[spt]
}

const (
	MinType = MajorTypeShuangXiu
	MaxType = MajorTypeFuQi
)
