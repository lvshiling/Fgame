package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/core/utils"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	cyclechargesinglemaxrewtemplate "fgame/fgame/game/welfare/cycle/charge_single_max_rew/template"
	cyclechargesinglemaxrewtypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家充值元宝
func playerChargeSingleCycle(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	chargeNum, ok := data.(int32)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeCycleCharge
	subType := welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRew

	//每日单笔充值
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.RefreshActivityData(typ, subType)

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*cyclechargesinglemaxrewtypes.CycleSingleChargeMaxRewInfo)
		if chargeNum > info.MaxSingleChargeNum {
			info.MaxSingleChargeNum = chargeNum
		}

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*cyclechargesinglemaxrewtemplate.GroupTemplateCycleSingleMaxRew)
		descTempList := groupTemp.GetCurDayTempDescList(info.CycleDay)
		for _, temp := range descTempList {
			needGold := temp.Value2
			if chargeNum < needGold {
				continue
			}

			if utils.ContainInt32(info.CanRewRecord, needGold) {
				continue
			}

			if utils.ContainInt32(info.ReceiveRewRecord, needGold) {
				continue
			}

			info.AddCanRewRecord(needGold)
			break
		}

		welfareManager.UpdateObj(obj)

		scMsg := pbutil.BuildSCOpenActivityCycleSingleChargeMaxRewInfoNotice(groupId, info.MaxSingleChargeNum, info.CanRewRecord)
		pl.SendMsg(scMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeSingleCycle))
}
