package common

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/battle/battle"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//排行数据变化
func rankValueChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}

	s := p.GetScene()
	if s == nil {
		return
	}

	eventData, ok := data.(*battle.BattlePlayerActivityRankDataChangedEventData)
	if !ok {
		return
	}

	rankType := eventData.GetRankType()
	rankData := eventData.GetRankData()

	//更新场景活动数据
	s.UpdatePlayer(rankType, p.GetId(), p.GetName(), rankData.GetRankValue(rankType))

	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityRankDataChanged, event.EventListenerFunc(rankValueChanged))
}
