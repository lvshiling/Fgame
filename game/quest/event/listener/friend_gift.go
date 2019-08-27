package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//送花
func friendGift(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeSendFlower, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendGift, event.EventListenerFunc(friendGift))
}
