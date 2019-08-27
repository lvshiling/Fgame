package listener

import (
	"fgame/fgame/core/event"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/battle/battle"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家比武大会积分达到x
func playerArenaPVPJiFenChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*battle.BattlePlayerActivityRankDataChangedEventData)
	if !ok {
		return
	}

	activityType := eventData.GetRankData().GetActivityType()
	if activityType != activitytypes.ActivityTypeArenapvp {
		return
	}

	rankData := eventData.GetRankData()
	rankType := eventData.GetRankType()
	jiFen := rankData.GetRankValue(rankType)

	err = questlogic.SetQuestDataSurpass(pl, questtypes.QuestSubTypeArenaPVPJiFen, 0, int32(jiFen))

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerActivityRankDataChanged, event.EventListenerFunc(playerArenaPVPJiFenChange))
}
