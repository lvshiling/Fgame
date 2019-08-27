package types

//沉迷状态
type WallowState int32

const (
	WallowStateNone WallowState = iota
	WallowStateThreeHour
	WallowStateFiveHour
)

var (
	wallowStateRateMap = map[WallowState]float64{
		WallowStateNone:      1,
		WallowStateThreeHour: 0.5,
		WallowStateFiveHour:  0,
	}
)

func (ws WallowState) Rate() float64 {
	return wallowStateRateMap[ws]
}
