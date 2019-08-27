package player

import (
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	hallzhuanshengtypes "fgame/fgame/game/welfare/hall/zhuansheng/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 转生冲刺
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeZhaunSheng, playerwelfare.ActivityObjInfoInitFunc(systemActivateZhuanShengInitInfo))
}

func systemActivateZhuanShengInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*hallzhuanshengtypes.ZhuanShengInfo)

	info.RewRecord = []int32{}
	info.IsMail = false

	propertyManager := obj.GetPlayer().GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	initZhuanSheng := propertyManager.GetZhuanSheng()
	info.ZhuanSheng = initZhuanSheng
}
