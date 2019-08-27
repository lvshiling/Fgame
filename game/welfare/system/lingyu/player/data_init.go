package player

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	systemlingyutypes "fgame/fgame/game/welfare/system/lingyu/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 领域系统激活
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeSystemActivate, welfaretypes.OpenActivitySystemActivateSubTypeLingYu, playerwelfare.ActivityObjInfoInitFunc(systemActivateLingYuInitInfo))
}

func systemActivateLingYuInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*systemlingyutypes.SystemLingYuInfo)
	info.StartTime = int64(0)
	info.IsOpen = false
	info.IsActivate = false
}
