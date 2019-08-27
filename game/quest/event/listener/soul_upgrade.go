package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	souleventtypes "fgame/fgame/game/soul/event/types"
)

//帝魂魂技升级
func soulUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*souleventtypes.SoulUpgradeEventData)
	if !ok {
		return
	}

	soulTag := eventData.GetSoulTag()
	orderLevel := eventData.GetNewOrder()
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSoulUpgradeLevel, int32(soulTag), orderLevel)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulUpgrade, event.EventListenerFunc(soulUpgrade))
}
