package types

const (
	TuLongRankSize = 10
)

type EggStatusType int32

const (
	//龙蛋
	EggStatusTypeInit EggStatusType = iota
	//boss
	EggStatusTypeBoss
	//已击杀
	EggStatusTypeDead
)
