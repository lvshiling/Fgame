package player

import (
	alliancecheertypes "fgame/fgame/game/welfare/alliance/cheer/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 城战助威
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeAlliance, playerwelfare.ActivityObjInfoInitFunc(allianceCheerInitInfo))
}

func allianceCheerInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*alliancecheertypes.AllianceCheerInfo)
	info.CheerGoldNum = 0
	info.IsEmail = false
}
