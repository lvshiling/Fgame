package types

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


//排行榜
const (
	ShenMoRankSize = 999
)
