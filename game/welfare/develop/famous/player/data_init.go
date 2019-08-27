package player

import (
	developfamoustypes "fgame/fgame/game/welfare/develop/famous/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 限时折扣
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeDevelop, welfaretypes.OpenActivityDefaultSubTypeDefault, playerwelfare.ActivityObjInfoInitFunc(developFameInitInfo))
}

func developFameInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*developfamoustypes.DevelopFameInfo)
	info.FavorableNum = 0
	info.FeedTimesMap = make(map[int32]int32)
	info.RewRecord = []int32{}
	info.IsEmail = false
}
