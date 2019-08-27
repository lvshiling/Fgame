package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	playercharge "fgame/fgame/game/charge/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	drewchargedrewtemplate "fgame/fgame/game/welfare/drew/charge_drew/template"
	drewchargedrewtypes "fgame/fgame/game/welfare/drew/charge_drew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

//玩家充值元宝
func playerChargeDrew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeChargeDrew
	//充值抽奖活动
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	drewTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range drewTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*drewchargedrewtypes.LuckyChargeDrewInfo)

		chargeManager := obj.GetPlayer().GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)

		isStartDay, _ := timeutils.IsSameDay(obj.GetStartTime(), now)

		//特殊处理
		//第一天同步数据
		if isStartDay {
			if info.GoldNum != int32(chargeManager.GetTodayChargeNum()) {
				addNum := int32(chargeManager.GetTodayChargeNum()) - info.GoldNum
				if addNum > 0 {
					info.GoldNum += addNum
					info.LeftConvertNum += addNum
				}
			}
		} else {
			info.GoldNum += goldNum
			info.LeftConvertNum += goldNum
		}

		groupInteface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInteface == nil {
			continue
		}
		groupTemp := groupInteface.(*drewchargedrewtemplate.GroupTemplateChargeDrew)
		convertRate := groupTemp.GetChargeDrewConvertRate()
		convertLimit := groupTemp.GetChargeDrewConvertLimit()
		minCycle := groupTemp.GetChargeDrewMinCycleTimes()
		info.CountLeftTimes(convertLimit, convertRate, minCycle)
		welfareManager.UpdateObj(obj)

		scMsg := pbutil.BuildSCOpenActivityDrewTimesNotice(groupId, info.LeftTimes)
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeDrew))
}
