package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	shopdiscounttemplate "fgame/fgame/game/welfare/shopdiscount/template"
	shopdiscounttypes "fgame/fgame/game/welfare/shopdiscount/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家充值
func playerChargeShopDiscount(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	chargeNum, ok := data.(int32)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeShopDiscount
	subType := welfaretypes.OpenActivityDefaultSubTypeDefault

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*shopdiscounttypes.ShopDiscountInfo)
		info.PeriodChargeNum += chargeNum
		welfareManager.UpdateObj(obj)

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*shopdiscounttemplate.GroupTemplateShopDiscount)
		needChargeNum := groupTemp.GetShopDiscountNeedChargeNum()
		if info.PeriodChargeNum < needChargeNum {
			continue
		}

		data := welfareeventtypes.CreatePlayerShopDiscountEventData(groupId, groupTemp.GetShopDiscountType(), obj.GetStartTime(), obj.GetEndTime())
		gameevent.Emit(welfareeventtypes.EventTypeShopDiscountActivite, pl, data)

		scMsg := pbutil.BuildSCOpenActivityPeriodChargeNotice(groupId, info.PeriodChargeNum)
		pl.SendMsg(scMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeShopDiscount))
}
