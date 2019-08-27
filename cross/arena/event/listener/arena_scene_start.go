package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	arenalogic "fgame/fgame/cross/arena/logic"
	"fgame/fgame/cross/arena/pbutil"
	arenascene "fgame/fgame/cross/arena/scene"
	gameevent "fgame/fgame/game/event"
)

//竞技场场景开始
func arenaSceneStart(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(arenascene.ArenaSceneData)
	scArenaSceneStart := pbutil.BuildSCArenaSceneStart()
	arenalogic.BroadcastArenaTeam(sd.GetTeam1(), scArenaSceneStart)
	arenalogic.BroadcastArenaTeam(sd.GetTeam2(), scArenaSceneStart)

	//TODO 机器人开始
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaSceneStart, event.EventListenerFunc(arenaSceneStart))
}
