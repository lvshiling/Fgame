package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	materialeventtypes "fgame/fgame/game/material/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//更新怪物波数
func playerMaterialRefreshGroup(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*materialeventtypes.RefreshGroupEventData)
	group := eventData.GetGroup()
	typ := eventData.GetMaterialType()
	err = questlogic.SetQuestDataSurpass(pl, questtypes.QuestSubTypeFuBenMonsterGroup, int32(typ), group)
	return
}

func init() {
	gameevent.AddEventListener(materialeventtypes.EventTypeMaterialRefreshGroup, event.EventListenerFunc(playerMaterialRefreshGroup))
}
