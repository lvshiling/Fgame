package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"
	"fgame/fgame/game/onearena/onearena"
	"fgame/fgame/game/onearena/pbutil"
	"fgame/fgame/game/player"
)

//灵池抢夺结束
func oneArenaRobEnd(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*onearenaeventtypes.OneArenaRobEndEventData)
	if !ok {
		return
	}
	sucess := eventData.GetSucess()
	levelType := eventData.GetLevel()
	pos := eventData.GetPos()
	ownerName := eventData.GetOwnerName()

	if sucess {
		onearena.GetOneArenaService().OneArenaRobSucess(pl, levelType, pos)
	} else {
		onearena.GetOneArenaService().OneArenaRobFail(pl, levelType, pos)
	}

	scOneArenaRobResult := pbutil.BuildSCOneArenaRobResult(sucess, ownerName, int32(levelType))
	pl.SendMsg(scOneArenaRobResult)
	return
}

func init() {
	gameevent.AddEventListener(onearenaeventtypes.EventTypeOneArenaRobEnd, event.EventListenerFunc(oneArenaRobEnd))
}
