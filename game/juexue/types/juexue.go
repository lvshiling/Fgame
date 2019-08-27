package types

type JueXueType int32

const (
	//三世陨
	JueXueTypeSanJie JueXueType = 1 + iota
	//六道灭
	JueXueTypeLiuDao
	//九霄碎
	JueXueTypeSui
)

func (jxt JueXueType) Valid() bool {
	switch jxt {
	case JueXueTypeSanJie,
		JueXueTypeLiuDao,
		JueXueTypeSui:
		return true
	}
	return false
}

func (t JueXueType) String() string {
	return jueXueMap[t]
}

var (
	jueXueMap = map[JueXueType]string{
		JueXueTypeSanJie: "突木桩",
		JueXueTypeLiuDao: "诛仙剑阵",
		JueXueTypeSui:    "六道傀儡",
	}
)

type JueXueStageType int32

const (
	//激活、强化
	JueXueStageTypeAorU JueXueStageType = 1 + iota
	//顿悟
	JueXueStageTypeInsight
)

func (jxst JueXueStageType) Valid() bool {
	switch jxst {
	case JueXueStageTypeAorU,
		JueXueStageTypeInsight:
		return true
	default:
		return false
	}
}
