package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//赛龙舟战力
type BoatRaceForceInfo struct {
	StartForce int64 `json:"startForce"` //活动开始时战力
	MaxForce   int64 `json:"maxForce"`   //活动期间最高战力
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeBoatRace, welfaretypes.OpenActivityDefaultSubTypeDefault, (*BoatRaceForceInfo)(nil))
}
