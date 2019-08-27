package player

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	xiuxianbooktypes "fgame/fgame/game/welfare/xiuxianbook/types"
)

func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipStrength, playerwelfare.ActivityObjInfoInitFunc(xiuxianBookInitInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipOpenLight, playerwelfare.ActivityObjInfoInitFunc(xiuxianBookInitInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipUpStar, playerwelfare.ActivityObjInfoInitFunc(xiuxianBookInitInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeLingTong, playerwelfare.ActivityObjInfoInitFunc(xiuxianBookInitInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeDianXing, playerwelfare.ActivityObjInfoInitFunc(xiuxianBookInitInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeShenQi, playerwelfare.ActivityObjInfoInitFunc(xiuxianBookInitInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillXinFa, playerwelfare.ActivityObjInfoInitFunc(xiuxianBookInitInfo))
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillDiHun, playerwelfare.ActivityObjInfoInitFunc(xiuxianBookInitInfo))
}

func xiuxianBookInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info, _ := obj.GetActivityData().(*xiuxianbooktypes.XiuxianBookInfo)
	info.FirstTimeRewRecord = -1
	info.ChargeNum = 0
	info.HasReceiveRecord = []int32{}
	info.MaxLevel = 0
}
