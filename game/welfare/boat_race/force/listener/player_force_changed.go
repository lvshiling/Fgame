package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
	boatraceforcetypes "fgame/fgame/game/welfare/boat_race/force/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	ranklogic "fgame/fgame/game/welfare/rank/logic"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家战力变化
func playerForceChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s != nil && s.MapTemplate().GetMapType() == scenetypes.SceneTypeChengZhan {
		return
	}
	curForce := pl.GetForce()
	typ := welfaretypes.OpenActivityTypeBoatRace
	subType := welfaretypes.OpenActivityDefaultSubTypeDefault
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*boatraceforcetypes.BoatRaceForceInfo)

		if curForce < info.MaxForce {
			continue
		}
		diffForce := curForce - info.MaxForce
		if diffForce > 0 {
			info.MaxForce = curForce
			welfareManager.UpdateObj(obj)
			ranklogic.UpdateAddCountRankData(pl, groupId, int32(diffForce))
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerForceChanged, event.EventListenerFunc(playerForceChanged))
}
