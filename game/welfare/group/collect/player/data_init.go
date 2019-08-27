package player

import (
	groupcollectenum "fgame/fgame/game/welfare/group/collect/enum"
	groupcollecttypes "fgame/fgame/game/welfare/group/collect/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 卡牌收集
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeGroup, welfaretypes.OpenActivityGroupSubTypeCollectPoker, playerwelfare.ActivityObjInfoInitFunc(collectRewInitInfo))
}

func collectRewInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*groupcollecttypes.CollectRewInfo)
	info.HadPokerList = []int32{}
	info.NonePokerList = groupcollectenum.GetInitPokerList()
	info.RewRecord = []groupcollectenum.PokerType{}
}
