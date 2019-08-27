package types

type SanGongType int32

const (
	SanGongTypeJunior SanGongType = iota //初级散功
	SanGongTypeMiddle                    //中级散功
	SanGongTypeSenior                    //高级散功
)

func (t SanGongType) Valid() bool {
	switch t {
	case SanGongTypeJunior,
		SanGongTypeMiddle,
		SanGongTypeSenior:
		return true
	default:
		return false
	}
}

var (
	sanGongMap = map[SanGongType]int32{
		SanGongTypeJunior: 1,
		SanGongTypeMiddle: 5,
		SanGongTypeSenior: 10,
	}
)

// 散功等级
func (t SanGongType) GetReduceLevelNum() int32 {
	return sanGongMap[t]
}

//潜能类型
type QianNengType int32

const (
	QianNengTypeTiZhi QianNengType = iota + 1 //体质
	QianNengTypeLiDao                         //力道
	QianNengTypeJinGu                         //筋骨
)

func (t QianNengType) Valid() bool {
	switch t {
	case QianNengTypeTiZhi,
		QianNengTypeLiDao,
		QianNengTypeJinGu:
		return true
	default:
		return false
	}
}
