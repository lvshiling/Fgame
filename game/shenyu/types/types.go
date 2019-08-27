package types

type ShenYuSceneRankType int32

func (t ShenYuSceneRankType) GetRankType() int32 {
	return int32(t)
}

const (
	ShenYuSceneRankTypeKey ShenYuSceneRankType = iota //神域排行榜
)
