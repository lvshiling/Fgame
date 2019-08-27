package player

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	tongtiantatypes "fgame/fgame/game/welfare/tongtianta/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 通天塔-宝宝
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeLingTong, playerwelfare.ActivityObjInfoInitFunc(initInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeMingGe, playerwelfare.ActivityObjInfoInitFunc(initInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeTuLong, playerwelfare.ActivityObjInfoInitFunc(initInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeShengHen, playerwelfare.ActivityObjInfoInitFunc(initInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeZhenFa, playerwelfare.ActivityObjInfoInitFunc(initInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeBaby, playerwelfare.ActivityObjInfoInitFunc(initInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeTongTianTa, welfaretypes.TongTianTaSubTypeDianXing, playerwelfare.ActivityObjInfoInitFunc(initInfo))

}

func initInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*tongtiantatypes.TongTianTaInfo)
	info.Record = make([]int32, 0, 8)
	info.MinForce = -1
	info.MaxForce = -1
	info.IsEmail = false
	info.ChargeNum = 0
}
