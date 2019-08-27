package types

type ShengTanSceneRankType int32

func (t ShengTanSceneRankType) GetRankType() int32 {
	return int32(t)
}

const (
	ShengTanSceneRankTypeDamage ShengTanSceneRankType = iota
)
