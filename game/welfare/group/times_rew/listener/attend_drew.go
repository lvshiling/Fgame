package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	grouptimesrewtypes "fgame/fgame/game/welfare/group/times_rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家活动抽奖
func playerAttendDrew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*welfareeventtypes.PlayerAttendDrewEventData)
	if !ok {
		return
	}

	//更新次数奖励
	updateDrewTimesRewData(pl, eventData.GetGroupId(), eventData.GetAttendNum())

	return
}

func updateDrewTimesRewData(pl player.Player, attendGroupId, attendNum int32) {
	typ := welfaretypes.OpenActivityTypeGroup
	subType := welfaretypes.OpenActivityGroupSubTypeTimesRew
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		if !timeTemp.IsRelationToGroup(attendGroupId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*grouptimesrewtypes.TimesRewInfo)
		info.Times += attendNum
		welfareManager.UpdateObj(obj)
	}
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeAttendDrew, event.EventListenerFunc(playerAttendDrew))
}
