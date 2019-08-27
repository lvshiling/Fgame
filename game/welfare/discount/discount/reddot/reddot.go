package reddot

// import (
// 	"fgame/fgame/game/player"
// 	playertypes "fgame/fgame/game/player/types"
// 	playerproperty "fgame/fgame/game/property/player"
// 	"fgame/fgame/game/reddot/reddot"
// 	discountdiscounttypes "fgame/fgame/game/welfare/discount/discount/types"
// 	welfarelogic "fgame/fgame/game/welfare/logic"
// 	playerwelfare "fgame/fgame/game/welfare/player"
// 	welfaretemplate "fgame/fgame/game/welfare/template"
// 	welfaretypes "fgame/fgame/game/welfare/types"
// )

// func init() {
// 	reddot.Register(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeCommon, reddot.HandlerFunc(handleRedDotDiscount))
// }

// //限时礼包红点
// func handleRedDotDiscount(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
// 	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
// 	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
// 	if !welfarelogic.IsOnActivityTime(groupId) {
// 		return
// 	}

// 	buyRecord := make(map[int32]int32)
// 	obj := welfareManager.GetOpenActivity(groupId)
// 	if obj != nil {
// 		info := obj.GetActivityData().(*discountdiscounttypes.DiscountInfo)
// 		buyRecord = info.BuyRecord
// 	}

// 	curGold := propertyManager.GetGold()
// 	discountDay := welfarelogic.CountCurActivityDay(groupId)
// 	discountTempList := welfaretemplate.GetWelfareTemplateService().GetDiscountTemplateByDayGroup(discountDay)
// 	for _, temp := range discountTempList {
// 		if temp.LimitCount > 0 {
// 			buyTimes := buyRecord[temp.Index]
// 			if buyTimes >= temp.LimitCount {
// 				continue
// 			}
// 		}

// 		if curGold < int64(temp.UseGold) {
// 			continue
// 		}

// 		isNotice = true
// 		return
// 	}
// 	return
// }
