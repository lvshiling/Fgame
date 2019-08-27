package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	hallzhuanshengtypes "fgame/fgame/game/welfare/hall/zhuansheng/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家转生变化
func playerZhuanShengChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	zhuanShu := data.(int32)

	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeZhaunSheng
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		//刷新
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*hallzhuanshengtypes.ZhuanShengInfo)
		info.UpdateZhaunSheng(zhuanShu)
		welfareManager.UpdateObj(obj)
	}

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerZhuanShengChanged, event.EventListenerFunc(playerZhuanShengChanged))
}
