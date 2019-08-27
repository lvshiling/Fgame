package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	drewcrazyboxtemplate "fgame/fgame/game/welfare/drew/crazy_box/template"
	drewcrazyboxtypes "fgame/fgame/game/welfare/drew/crazy_box/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家消耗元宝
func playerCostGoldCrazyBox(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int64)
	if !ok {
		return
	}
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeCrazyBox

	//疯狂宝箱
	welfareManager := p.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
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
		groupTemp := groupInterface.(*drewcrazyboxtemplate.GroupTemplateCrazyBox)

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*drewcrazyboxtypes.CrazyBoxInfo)
		oldTotalTimes := groupTemp.GetCrazyBoxTotalTimes(info.GoldNum)
		info.GoldNum += int32(goldNum)
		welfareManager.UpdateObj(obj)
		newTotalTimes := groupTemp.GetCrazyBoxTotalTimes(info.GoldNum)
		if newTotalTimes > oldTotalTimes {
			drewTimes := newTotalTimes - info.AttendTimes
			scMsg := pbutil.BuildSCOpenActivityCrazyBoxTimesNotice(groupId, drewTimes)
			p.SendMsg(scMsg)
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerGoldCost, event.EventListenerFunc(playerCostGoldCrazyBox))
}
