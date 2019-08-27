package types

//星数
type SoulRuinsStarNumType int32

const (
	//0颗星
	SoulRuinsStarNumTypeZero SoulRuinsStarNumType = iota
	//1颗星
	SoulRuinsStarNumTypeOne
	//2颗星
	SoulRuinsStarNumTypeTwo
	//3颗星
	SoulRuinsStarNumTypeThree
)

const (
	//通关的最小星数
	MinStar = int32(SoulRuinsStarNumTypeOne)
	//通关的最大星数
	MaxStar = int32(SoulRuinsStarNumTypeThree)
)

//难度
type SoulRuinsType int32

const (
	//普通
	SoulRuinsTypeEasy SoulRuinsType = 1 + iota
	//困难
	SoulRuinsTypeHard
)

func (srt SoulRuinsType) Valid() bool {
	switch srt {
	case SoulRuinsTypeEasy,
		SoulRuinsTypeHard:
		return true
	}
	return false
}

const (
	//难度最小值
	SoulRuinsTypeMin = SoulRuinsTypeEasy
	//难度最大值
	SoulRuinsTypeMax = SoulRuinsTypeHard
)

//事件类型
type SoulRuinsEventType int32

const (
	//不触发
	SoulRuinsEventTypeNot SoulRuinsEventType = iota
	//boss
	SoulRuinsEventTypeBoss
	//帝魂降临
	SoulRuinsEventTypeSoul
	//马贼
	SoulRuinsEventTypeRobber
)

func (sret SoulRuinsEventType) Valid() bool {
	switch sret {
	case SoulRuinsEventTypeBoss,
		SoulRuinsEventTypeSoul,
		SoulRuinsEventTypeRobber:
		return true
	}
	return false
}

//阶段类型
type SoulRuinsStageType int32

const (
	//杀怪
	SoulRuinsStageTypeKillMonster SoulRuinsStageType = 1 + iota
	//触发事件
	SoulRuinsStageTypeEvent
	//第二阶段
	SoulRuinsStageTypeSecond
	//完成
	SoulRuinsStageTypeFinshed
)

func (srst SoulRuinsStageType) Valid() bool {
	switch srst {
	case SoulRuinsStageTypeKillMonster,
		SoulRuinsStageTypeEvent,
		SoulRuinsStageTypeSecond,
		SoulRuinsStageTypeFinshed:
		return true
	}
	return false
}
