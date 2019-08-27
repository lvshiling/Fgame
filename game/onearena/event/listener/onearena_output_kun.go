package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"
	onearenalogic "fgame/fgame/game/onearena/logic"
	"fgame/fgame/game/onearena/onearena"
)

//灵池产出鲲
func oneArenaOutputKun(target event.EventTarget, data event.EventData) (err error) {
	oneArenaObj, ok := target.(*onearena.OneArenaObject)
	if !ok {
		return
	}
	num, ok := data.(int32)
	if !ok {
		return
	}
	ownerId := oneArenaObj.OwnerId
	level := oneArenaObj.Level
	pos := oneArenaObj.Pos
	err = onearenalogic.OneArenaOutputKun(ownerId, level, pos, num)
	return
}

func init() {
	gameevent.AddEventListener(onearenaeventtypes.EventTypeOneArenaOutputKun, event.EventListenerFunc(oneArenaOutputKun))
}
