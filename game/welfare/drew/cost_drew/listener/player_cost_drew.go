package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	drewcostdrewtemplate "fgame/fgame/game/welfare/drew/cost_drew/template"
	drewcostdrewtypes "fgame/fgame/game/welfare/drew/cost_drew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家消费元宝
func playerGoldCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int64)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeCostDrew
	//消费抽奖活动
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	drewTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range drewTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*drewcostdrewtypes.LuckyCostDrewInfo)
		info.GoldNum += int32(goldNum)
		info.LeftConvertNum += int32(goldNum)

		groupInteface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInteface == nil {
			continue
		}
		groupTemp := groupInteface.(*drewcostdrewtemplate.GroupTemplateCostDrew)
		convertRate := groupTemp.GetCostDrewConvertRate()
		convertLimit := groupTemp.GetCostDrewConvertLimit()
		minCycle := groupTemp.GetCostDrewMinCycleTimes()
		info.CountLeftTimes(convertLimit, convertRate, minCycle)
		welfareManager.UpdateObj(obj)

		scMsg := pbutil.BuildSCOpenActivityDrewTimesNotice(groupId, info.LeftTimes)
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerGoldCost, event.EventListenerFunc(playerGoldCost))
}
