package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	souleventtypes "fgame/fgame/game/soul/event/types"
)

//帝魂镶嵌
func soulEmbed(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*souleventtypes.SoulEmbedEventData)
	if !ok {
		return
	}
	soulTag := eventData.GetNewTag()
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSoulSpecialEmbed, int32(soulTag), 1)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulEmbed, event.EventListenerFunc(soulEmbed))
}
