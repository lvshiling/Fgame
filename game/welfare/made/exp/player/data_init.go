package player

import (
	madeexptypes "fgame/fgame/game/welfare/made/exp/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 炼制-经验
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeMade, welfaretypes.OpenActivityMadeSubTypeResource, playerwelfare.ActivityObjInfoInitFunc(madeResInitInfo))
}

func madeResInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*madeexptypes.MadeInfo)
	info.Times = 0
}
