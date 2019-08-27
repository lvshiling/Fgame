package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"

	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	tongtiantatypes "fgame/fgame/game/welfare/tongtianta/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, event.EventListenerFunc(forceChange))
}

func forceChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	propertyEffectorType := data.(playerpropertytypes.PropertyEffectorType)
	subType, exist := tongtiantatypes.PlayerPropertyEffectTypeToTongTianTaSubType(propertyEffectorType)
	if !exist {
		return
	}
	effectTypeList, ok := tongtiantatypes.TongTianTaSubTypeTOPlayerPropertyEffectType(subType)
	if !ok {
		return
	}
	power := int32(0)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	for _, effType := range effectTypeList {
		power += int32(propertyManager.GetModuleForce(effType))
	}

	typ := welfaretypes.OpenActivityTypeTongTianTa

	welfareManager := pl.GetPlayerDataManager(types.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityData(typ, subType)
	if err != nil {
		return
	}

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, tiemTemp := range timeTempList {
		groupId := tiemTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*tongtiantatypes.TongTianTaInfo)

		if info.MinForce == -1 {
			info.MinForce = power
			info.MaxForce = power
			welfareManager.UpdateObj(obj)
			continue
		}

		// 战斗力下降的不能领取
		if power <= info.MaxForce {
			continue
		}

		// 模块战斗力增加
		info.MaxForce = power

		welfareManager.UpdateObj(obj)
	}
	return
}
