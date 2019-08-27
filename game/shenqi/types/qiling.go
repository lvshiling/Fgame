package types

type QiLingType int32

const (
	//普通器灵
	QiLingTypeNormal QiLingType = iota
	//超级器灵
	QiLingTypeSuper
)

func (spt QiLingType) Valid() bool {
	switch spt {
	case QiLingTypeNormal,
		QiLingTypeSuper:
		return true
	}
	return false
}

type QiLingSubType interface {
	SubType() int32
	Valid() bool
	String() string
}

type QiLingSubTypeFactory interface {
	CreateQiLingSubType(subType int32) QiLingSubType
}

type QiLingSubTypeFactoryFunc func(subType int32) QiLingSubType

func (t QiLingSubTypeFactoryFunc) CreateQiLingSubType(subType int32) QiLingSubType {
	return t(subType)
}

var (
	qiLingSubTypeFactoryMap = make(map[QiLingType]QiLingSubTypeFactory)
	QiLingSubTypeStringMap  = make(map[QiLingType]map[QiLingSubType]string)
)

func init() {
	qiLingSubTypeFactoryMap[QiLingTypeNormal] = QiLingSubTypeFactoryFunc(CreateQiLingNormalSubType)
	qiLingSubTypeFactoryMap[QiLingTypeSuper] = QiLingSubTypeFactoryFunc(CreateQiLingSuperSubType)

	QiLingSubTypeStringMap[QiLingTypeNormal] = qiLingNormalSubTypeStringMap
	QiLingSubTypeStringMap[QiLingTypeSuper] = qiLingSuperSubTypeStringMap
}

func CreateQiLingSubType(typ QiLingType, subType int32) QiLingSubType {
	factory, ok := qiLingSubTypeFactoryMap[typ]
	if !ok {
		panic("shenqi:CreateQiLingSubType 应该是ok的")
	}
	return factory.CreateQiLingSubType(subType)
}

type QiLingNormalSubType int32

const (
	//开
	QiLingNormalSubTypeKai QiLingNormalSubType = iota
	//休
	QiLingNormalSubTypeXiu
	//生
	QiLingNormalSubTypeSheng
	//伤
	QiLingNormalSubTypeShang
	//杜
	QiLingNormalSubTypeDu
	//景
	QiLingNormalSubTypeJing
	//死
	QiLingNormalSubTypeSi
	//惊
	QiLingNormalSubTypeJin
)

func (spt QiLingNormalSubType) SubType() int32 {
	return int32(spt)
}

func (spt QiLingNormalSubType) Valid() bool {
	switch spt {
	case QiLingNormalSubTypeKai,
		QiLingNormalSubTypeXiu,
		QiLingNormalSubTypeSheng,
		QiLingNormalSubTypeShang,
		QiLingNormalSubTypeDu,
		QiLingNormalSubTypeJing,
		QiLingNormalSubTypeSi,
		QiLingNormalSubTypeJin:
		return true
	}
	return false
}

func (spt QiLingNormalSubType) String() string {
	return QiLingSubTypeStringMap[QiLingTypeNormal][spt]
}

var (
	qiLingNormalSubTypeStringMap = map[QiLingSubType]string{
		QiLingNormalSubTypeKai:   "开",
		QiLingNormalSubTypeXiu:   "休",
		QiLingNormalSubTypeSheng: "生",
		QiLingNormalSubTypeShang: "伤",
		QiLingNormalSubTypeDu:    "杜",
		QiLingNormalSubTypeJing:  "景",
		QiLingNormalSubTypeSi:    "死",
		QiLingNormalSubTypeJin:   "惊",
	}
)

func CreateQiLingNormalSubType(subType int32) QiLingSubType {
	return QiLingNormalSubType(subType)
}

type QiLingSuperSubType int32

const (
	//天
	QiLingSuperSubTypeTian QiLingSuperSubType = iota
	//地
	QiLingSuperSubTypeDi
)

func (spt QiLingSuperSubType) SubType() int32 {
	return int32(spt)
}

func (spt QiLingSuperSubType) Valid() bool {
	switch spt {
	case QiLingSuperSubTypeTian,
		QiLingSuperSubTypeDi:
		return true
	}
	return false
}

func (spt QiLingSuperSubType) String() string {
	return QiLingSubTypeStringMap[QiLingTypeSuper][spt]
}

var (
	qiLingSuperSubTypeStringMap = map[QiLingSubType]string{
		QiLingSuperSubTypeTian: "天",
		QiLingSuperSubTypeDi:   "地",
	}
)

func CreateQiLingSuperSubType(subType int32) QiLingSubType {
	return QiLingSuperSubType(subType)
}
