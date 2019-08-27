package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	playerrealm "fgame/fgame/game/realm/player"
	hallrealmtypes "fgame/fgame/game/welfare/hall/realm/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//天劫塔挑战结果
func realmChallengeResult(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if pl == nil {
		return
	}
	sucessful, ok := data.(bool)
	if !ok {
		return
	}
	if !sucessful {
		return
	}
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	realmManager := pl.GetPlayerDataManager(playertypes.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeRealm

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		level := realmManager.GetTianJieTaLevel()
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*hallrealmtypes.WelfareRealmChallengeInfo)
		if info.Level >= level {
			continue
		}
		info.Level = level
		welfareManager.UpdateObj(obj)

	}

	return
}

func init() {
	gameevent.AddEventListener(realmeventtypes.EventTypeRealmResult, event.EventListenerFunc(realmChallengeResult))
}
