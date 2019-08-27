package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	xiuxianbookeventtypes "fgame/fgame/game/welfare/xiuxianbook/event/types"
)

func init() {
	gameevent.AddEventListener(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, event.EventListenerFunc(xiuxianbookreUpdateObj))
}

func xiuxianbookreUpdateObj(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	typ := welfaretypes.OpenActivityTypeXiuxianBook

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	for subType := welfaretypes.MinOpenActivityXiuxianBookSubType; subType <= welfaretypes.MaxOpenActivityXiuxianBookSubType; subType++ {
		err = welfareManager.RefreshActivityData(typ, subType)
		if err != nil {
			continue
		}
		timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
		for _, timeTemp := range timeTempList {
			groupId := timeTemp.Group
			if !welfarelogic.IsOnActivityTime(groupId) {
				continue
			}
			obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
			welfareManager.UpdateObj(obj)
		}
	}
	return
}
