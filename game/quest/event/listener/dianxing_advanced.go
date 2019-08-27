package listener

import (
	"fgame/fgame/core/event"
	dianxingeventtypes "fgame/fgame/game/dianxing/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家点星系统进阶
func playerDianXingAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId := data.(int32)
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeDianXing), advancedId)
	return
}

func init() {
	gameevent.AddEventListener(dianxingeventtypes.EventTypeDianXingAdvanced, event.EventListenerFunc(playerDianXingAdavanced))
}
