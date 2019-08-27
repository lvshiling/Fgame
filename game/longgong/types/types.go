package types

type LongGongSceneRankType int32

func (t LongGongSceneRankType) GetRankType() int32 {
	return int32(t)
}

const (
	LongGongSceneRankTypeDamage LongGongSceneRankType = iota
)

type HeiLongStatusType int32

const (
	//未出生
	HeiLongStatusTypeInit HeiLongStatusType = iota
	//boss
	HeiLongStatusTypeLive
	//已击杀
	HeiLongStatusTypeDead
)
