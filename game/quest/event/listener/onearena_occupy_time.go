package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家灵池占领时间
func playerOneArenaOccupyTime(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	// oneArenaObj, ok := data.(*playeronearena.PlayerOneArenaObject)
	// if !ok {
	// 	return
	// }
	// now := global.GetGame().GetTimeService().Now()
	// level := int32(oneArenaObj.Level)
	// occupyTime := now - oneArenaObj.RobTime
	err = questlogic.SetQuestEmbedData(pl, questtypes.QuestSubTypeOneArenaOccupyTime)
	return
}

func init() {
	gameevent.AddEventListener(onearenaeventtypes.EventTypeOneArenaOccupyTime, event.EventListenerFunc(playerOneArenaOccupyTime))
}
