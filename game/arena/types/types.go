package types

type FourGodType int32

const (
	FourGodTypeQingLong FourGodType = iota
	FourGodTypeBaiHu
	FourGodTypeZhuQue
	FourGodTypeXuanWu
)

var (
	fourGodMap = map[FourGodType]string{
		FourGodTypeQingLong: "青龙",
		FourGodTypeBaiHu:    "白虎",
		FourGodTypeZhuQue:   "朱雀",
		FourGodTypeXuanWu:   "玄武",
	}
)

func (t FourGodType) Valid() bool {
	switch t {
	case FourGodTypeQingLong,
		FourGodTypeBaiHu,
		FourGodTypeZhuQue,
		FourGodTypeXuanWu:
		return true
	}
	return false
}

func (t FourGodType) String() string {
	return fourGodMap[t]

}

//
type ArenaType int32

const (
	ArenaTypeArena    = iota //连胜
	ArenaTypeQingLong        //废弃
	ArenaTypeBaiHu           //废弃
	ArenaTypeZhuQue          //废弃
	ArenaTypeXuanWu          //废弃
	ArenaTypeFail            //连败

)

func (t ArenaType) Valid() bool {
	switch t {
	case ArenaTypeArena,
		ArenaTypeQingLong,
		ArenaTypeBaiHu,
		ArenaTypeZhuQue,
		ArenaTypeXuanWu,
		ArenaTypeFail:
		return true
	}
	return false
}

//排行榜类型
type RankTimeType int32

const (
	//本周
	RankTimeTypeThis RankTimeType = iota
	//上周
	RankTimeTypeLast
)

func (t RankTimeType) Vaild() bool {
	switch t {
	case RankTimeTypeThis,
		RankTimeTypeLast:
		return true
	}
	return false
}
