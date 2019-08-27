package player

import (
	rewpoolstypes "fgame/fgame/game/welfare/drew/rew_pools/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeRewPools, playerwelfare.ActivityObjInfoInitFunc(rewPoolsInitInfo))
}

func rewPoolsInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info, _ := obj.GetActivityData().(*rewpoolstypes.RewPoolsInfo)
	info.Position = int32(0)
	info.BackTimes = int32(0)
	info.RecordTime = int64(0)
}
