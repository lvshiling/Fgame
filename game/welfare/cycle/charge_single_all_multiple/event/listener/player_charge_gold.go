package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	chargesingleallmultipletemplate "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/template"
	chargesingleallmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func playerChargeGold(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	// chargeNum, ok := data.(int32)
	// if !ok {
	// 	return
	// }
	typ := welfaretypes.OpenActivityTypeCycleCharge
	subType := welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeAllRew

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.RefreshActivityData(typ, subType)

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*chargesingleallmultipletemplate.GroupTemplateCycleSingleAllMultiple)

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*chargesingleallmultipletypes.CycleChargeSingleAllMultipleInfo)
		// info.AddSingleChargeRecord(chargeNum)

		// descTempList := groupTemp.GetCurDayTempDescList(info.CycleDay)
		// for _, temp := range descTempList {
		// 	needGold := temp.Value2
		// 	if chargeNum < needGold {
		// 		continue
		// 	}
		// 	// // 第一天走refresh，同步今日消费
		// 	// now := global.GetGame().GetTimeService().Now()
		// 	// diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		// 	// if diff > 0 {
		// 	// 	num := chargeNum / needGold
		// 	// 	info.AddCanRewRecord(needGold, num)
		// 	// }
		// 	num := chargeNum / needGold
		// 	info.AddCanRewRecord(needGold, num)
		// }

		welfareManager.UpdateObj(obj)
		canRewRecord := groupTemp.GetCanRewRecordMap(info.CycleDay, info.GetCanRewRecord())
		scMsg := pbutil.BuildSCOpenActivityCycleSingleChargeAllRewInfoNotice(groupId, info.SingleChargeRecord, canRewRecord)
		pl.SendMsg(scMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeGold))
}
